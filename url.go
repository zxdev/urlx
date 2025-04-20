package urlx

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

// Package URL provides a public suffix aware url parser based on data from
// https://publicsuffix.org/ merged with a custom urlx.dat list.
//
// A public suffix is one under which Internet users can directly register
// names. It is related to, but different from, a TLD (top level domain).
//
// "com" is a TLD (top level domain). Top level means it has no dots.
//
// "com" is also a public suffix. Amazon and Google have registered different
// siblings under that domain: "amazon.com" and "google.com".
//
// "au" is another TLD, again because it has no dots. But it's not "amazon.au".
// Instead, it's "amazon.com.au".
//
// "com.au" isn't an actual TLD, because it's not at the top level (it has
// dots). But it is an eTLD (effective TLD), because that's the branching point
// for domain name registrars.
//
// Another name for "an eTLD" is "a public suffix". Often, what's more of
// interest is the eTLD+1, or one more label than the public suffix, or Apex.
// For example, browsers partition read/write access to HTTP cookies according
// to the eTLD+1. Web pages served from "amazon.com.au" can't read cookies from
// "google.com.au", but web pages served from "maps.google.com" can share
// cookies from "www.google.com", so you don't have to sign into Google Maps
// separately from signing into Google Web Search. Note that all four of those
// domains have 3 labels and 2 dots. The first two domains are each an eTLD+1,
// the last two are not (but share the same eTLD+1: "google.com").
//
// All of these domains have the same eTLD+1 or Apex hostname:
//   - "www.books.amazon.co.uk"
//   - "books.amazon.co.uk"
//   - "amazon.co.uk"
//
// Specifically, the eTLD+1 is "amazon.co.uk", because the eTLD is "co.uk".
//
// There is no closed form algorithm to calculate the eTLD of a domain.
// Instead, the calculation is data driven. This package provides a
// pre-compiled snapshot of Mozilla's PSL (Public Suffix List) data at
// https://publicsuffix.org/ which is used for detecting the apex form
// when parsing the url. The list automatically refreshed every 72 hours
// according the the foundations website recommendation.
type URL struct {
	tld                           map[string]uint8 // tld reference
	Apex, Host, Port, Path, TLD   string           // url segment
	IP, IDNA                      bool             // form type flags
	Kind                          uint8            // icann, private, custom flag
	onlyIP, onlyHost, onlyApex    bool             // conditional toggle type flags
	noWWW, noPath, noPort, noIDNA bool             // conditional toggle segment flags
	puny                          *idna.Profile    // operational element
	idx                           int              // operational element
	seg, host                     []string         // operational element
	ip                            net.IP           // operational element
	err                           error            // operational element
}

// Parser is the urlx configurator that will automatically download and
// refresh and then apply the publicsuffix.org and the urlx.dat custom list.
//
// Kind flag reports whether the suffix is managed by the Internet Corporation
// for Assigned Names and Numbers a privately managed domain (and in practice,
// not a top level domain) or an unmanaged top level domain (and not explicitly
// mentioned in the publicsuffix.org list), or a custom domain. For example,
// "foo.org" and "foo.co.uk" are ICANN domains, "foo.dyndns.org" is a private
// domain and "zxdev.com" when added to urlx.dat is a custom top level domain.
func Parser(path *string) *URL {

	u := new(URL)

	// assurances

	if path == nil || len(*path) == 0 {
		var resource = "dat"
		if runtime.GOOS == "linux" {
			resource = "/var/urlx"
		}
		os.Mkdir(resource, 0644)
		path = &resource
	}

	u.tld = make(map[string]uint8)

	// puny converts the internationalized domain name âbc.com to xn--bc-oia.com
	// to support an all ascii typeset; configuration
	u.puny = idna.New(idna.MapForLookup(), idna.Transitional(true))

	// refresh the publicsuffix.org source when aged
	// and build the tld map reference for use

	var count int
	var pubsuffix = filepath.Join(*path, "ps.dat")
	var info, err = os.Stat(pubsuffix)
	if err != nil || info.ModTime().Before(time.Now().Add(-time.Hour*72)) {
		r, err := http.Get("https://publicsuffix.org/list/effective_tld_names.dat")
		if err == nil && r != nil && r.StatusCode == http.StatusOK {
			w, err := os.Create(pubsuffix)
			if err == nil {
				io.Copy(w, r.Body)
				w.Close()
			}
		}
	}

	f, err := os.Open(pubsuffix)
	if err == nil {
		var kind uint8
		var row string
		var scanner = bufio.NewScanner(f)
		for scanner.Scan() {
			row = strings.TrimSpace(scanner.Text())
			switch {
			case len(row) == 0: // empty
			case kind != 1 && strings.Contains(row, "BEGIN ICANN DOMAINS"): // detect flag change
				kind = 1 // 0b0001
			case kind != 2 && strings.Contains(row, "BEGIN PRIVATE DOMAINS"): // detect flag change
				kind = 2 // 0b0010
			case strings.HasPrefix(row, "//"): // comment
			default:
				// ignore the *. rules for simplicity
				row = strings.TrimPrefix(row, "*.")
				u.tld[row] = kind
				count++
			}
		}
		f.Close()
	}

	// add the urlx.dat resource file or generate an empty
	// custom resource file when it is missing

	f, err = os.Open(filepath.Join(*path, "urlx.dat"))
	if err == nil {

		var row string
		var scanner = bufio.NewScanner(f)
		for scanner.Scan() {
			row = strings.TrimSpace(scanner.Text())
			switch {
			case len(row) == 0:
			case strings.HasPrefix(row, "//"):
			case strings.HasPrefix(row, "#"):
			default:
				u.tld[row] = 4 // 0b0100
				count++
			}
		}

		f.Close()

	} else {

		// generate empty file
		w, _ := os.Create(filepath.Join(*path, "urlx.dat"))
		fmt.Fprintln(w, "# urlx custom tld list | ", time.Now().Format(time.RFC3339)[:19])
		w.Close()

	}

	return u
}

/*

	methods
		len, string, reset
		ApexIsHost

*/

// Len reports number of tld items
func (u *URL) Len() int {
	if u.tld == nil {
		u.tld = make(map[string]byte)
	}
	return len(u.tld)
}

// String reconstitues url based on flag settings
func (u *URL) String() (url string) {

	url = u.Host

	if len(u.Port) > 0 {
		if strings.Contains(u.Host, ":") {
			url = "[" + u.Host + "]" // ipv6
		}
		url += ":" + u.Port // ipv4, domain
	}

	if len(u.Path) > 0 {
		url += "/" + u.Path
	}

	return
}

// HostIsApex is a boolean comparison test
func (u *URL) HostIsApex() bool {
	return len(u.Apex) > 0 && len(u.Host) == len(u.Apex)
}

// reset
func (u *URL) reset() {
	u.Apex, u.Host, u.Port, u.Path, u.TLD = "", "", "", "", ""
	u.IP, u.IDNA = false, false
	u.Kind = 0
}

/*

	parser flag toggles
		noWWW, noIDNA, noPath, noPort
		onlyIP, onlyHost, onlyApex

*/

// NoWWW; default on
func (u *URL) NoWWW() *URL { u.noWWW = !u.noWWW; return u } // on

// NoIDNA; default off
func (u *URL) NoIDNA() *URL { u.noIDNA = !u.noIDNA; return u } // off

// NoPath; default off
func (u *URL) NoPath() *URL { u.noPath = !u.noPath; return u } // off

// NoPort; default off
func (u *URL) NoPort() *URL { u.noPort = !u.noPort; return u } // off

// OnlyIP toggle; default off
//
//	IP:   on|off
//	Apex: off|on
//	Host: off|on
func (u *URL) OnlyIP() *URL {
	u.onlyIP, u.onlyApex, u.onlyHost = !u.onlyIP, false, false
	return u
}

// OnlyHost toggle; default off
//
//	Host: on|off
//	IP:   off|on
//	Apex: off|off
func (u *URL) OnlyHost() *URL {
	u.onlyHost, u.onlyIP, u.onlyApex = !u.onlyHost, false, false
	return u
}

// OnlyHost toggle; default off
//
//	Apex: on|off
//	Host: off|on
//	IP:   off|on
func (u *URL) OnlyApex() *URL {
	u.onlyApex, u.onlyHost, u.onlyIP = !u.onlyApex, false, false
	return u
}

/*

	url parser

*/

// Parse the url into consitituate parts. The constituant elements of the url
// are parsed into (tld, apex, host, port, path) when a domain is parsed and
// when the host is an IPv4/6 then (host, port path) are the only elements
// with the IP flag set. When the domain is a converted internationalized form
// the IDNA flag will be set. The kind field reports the type of domain:
//
//	Kind 1 0b0001 : icann managed tld
//	Kind 2 0b0010 : publicsuffix private tld
//	Kind 4 0b0100 : urlx custom tld
func (u *URL) Parse(url *string) (ok bool) {

	// reset url segments and type flags
	u.reset()
	u.Host = *url

	// strip query fragment
	if u.idx = strings.Index(u.Host, "#"); u.idx > 0 {
		u.Host = u.Host[:u.idx]
	}

	// strip query segment
	if u.idx = strings.Index(u.Host, "?"); u.idx > 0 {
		u.Host = u.Host[:u.idx]
	}

	// strip schemes
	if len(u.Host) > 8 {
		if u.idx = strings.Index(u.Host, "://"); u.idx > -1 {
			u.Host = u.Host[u.idx+3:]
		}
	}

	// extract path segment; remove trailing slashes
	u.seg = strings.SplitN(u.Host, "/", 2)
	if len(u.seg) == 2 {
		u.Host = u.seg[0]
		if !u.noPath {
			u.Path = strings.TrimSuffix(u.seg[1], "/")
		}
	}

	// remove user:host segment
	u.seg = strings.Split(u.Host, "@")
	if len(u.seg) > 1 {
		u.Host = u.seg[1]
	}

	// extract port
	switch {

	// unported ipv6
	case strings.HasSuffix(u.Host, "]") && strings.HasPrefix(u.Host, "["):
		u.Host = u.Host[1 : len(u.Host)-2]

	// ipv6
	case strings.HasPrefix(u.Host, "[") && strings.Contains(u.Host, "]:"):
		u.seg = strings.Split(u.Host[1:], "]:")
		u.Host = u.seg[0]
		if !u.noPort {
			u.Port = u.seg[1]
		}

	// !ipv6; ipv4|domain
	case strings.Contains(u.Host, "."):
		u.seg = strings.SplitN(u.Host, ":", 2)
		if len(u.seg) == 2 {
			u.Host = u.seg[0]
			if !u.noPort {
				u.Port = u.seg[1]
			}
		}
	}

	// detect and validate ipv4/6
	u.ip = net.ParseIP(u.Host)
	if u.IP = u.ip != nil; u.IP {
		if u.ip.IsUnspecified() || u.ip.IsLoopback() || u.ip.IsPrivate() {
			u.reset()
			return
		}
	}

	// validate domain vs ip type settings
	if u.onlyHost && u.IP || u.onlyIP && !u.IP {
		u.reset()
		return
	}

	// standardize domain and type settings
	if !u.IP {

		// standardize host
		u.Host = strings.TrimSpace(u.Host)       // strip whitespace
		u.Host = strings.ToLower(u.Host)         // lowercase
		u.Host = strings.TrimSuffix(u.Host, ".") // clean cannonical

		// flag for idna|punycode domains
		if !u.noIDNA {
			u.Host, u.err = u.puny.ToASCII(u.Host)
			u.IDNA = u.err == nil && strings.HasPrefix(u.Host, "xn--")
		}

		// localize apex and tld
		u.host = strings.Split(u.Host, ".")
		for u.idx = range len(u.host) {
			u.TLD = strings.Join(u.host[u.idx:], ".")
			if u.Kind = u.tld[u.TLD]; u.Kind > 0 {
				if len(u.TLD) == len(u.Host) {
					u.reset()
					return
				}
				break
			}
			u.Apex = strings.Join(u.host[u.idx:], ".")
		}

		// remove www label
		//  exception for icann tld, private, and custom forms only
		if !u.noWWW && strings.HasPrefix(u.Host, "www.") {
			if strings.HasPrefix(u.Host, "www.") && len(u.TLD)+4 != len(u.Host) {
				u.Host = u.Host[4:]
			}
		}

		// manage apex
		if u.onlyApex {
			u.Host, u.Path, u.Port = u.Apex, "", ""
		}

	}

	// final validation check
	if len(u.Host) < 254 && strings.ContainsAny(u.Host, ".:") {
		return true
	}

	return
}
