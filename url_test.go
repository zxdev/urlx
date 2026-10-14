package urlx_test

import (
	"testing"

	"github.com/zxdev/urlx"
)

func TestURL(t *testing.T) {

	/*
		=== RUN   TestURL
		    url_test.go:90: tld: 9959
		    url_test.go:115: origional  14x33 success= false
		    url_test.go:116: flags     compare= false ip= false idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host=  port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  com success= false
		    url_test.go:116: flags     compare= false ip= false idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host=  port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  api.c success= false
		    url_test.go:116: flags     compare= false ip= false idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host=  port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  duckdns.org success= false
		    url_test.go:116: flags     compare= false ip= false idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host=  port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  vfs.cloud9.eu-west-1.amazonaws.com success= false
		    url_test.go:116: flags     compare= false ip= false idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host=  port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  exàmple.com success= true
		    url_test.go:116: flags     compare= true ip= false idna= true kind= icann
		    url_test.go:117: elements  tld= com apex= xn--exmple-jta.com host= xn--exmple-jta.com port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  http://www.example.com success= true
		    url_test.go:116: flags     compare= true ip= false idna= false kind= icann
		    url_test.go:117: elements  tld= com apex= example.com host= example.com port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  blog.example.com:80/path/page success= true
		    url_test.go:116: flags     compare= false ip= false idna= false kind= icann
		    url_test.go:117: elements  tld= com apex= example.com host= blog.example.com port= 80 path= path/page
		    url_test.go:118:
		    url_test.go:115: origional  click.api.example.com:8888/path/page success= true
		    url_test.go:116: flags     compare= false ip= false idna= false kind= icann
		    url_test.go:117: elements  tld= com apex= example.com host= click.api.example.com port= 8888 path= path/page
		    url_test.go:118:
		    url_test.go:115: origional  www.fr success= true
		    url_test.go:116: flags     compare= true ip= false idna= false kind= icann
		    url_test.go:117: elements  tld= fr apex= www.fr host= www.fr port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  www.co.uk success= true
		    url_test.go:116: flags     compare= true ip= false idna= false kind= icann
		    url_test.go:117: elements  tld= co.uk apex= www.co.uk host= www.co.uk port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  www.duckdns.org success= true
		    url_test.go:116: flags     compare= true ip= false idna= false kind= publicsuffix
		    url_test.go:117: elements  tld= duckdns.org apex= www.duckdns.org host= www.duckdns.org port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  one.www.co.uk success= true
		    url_test.go:116: flags     compare= false ip= false idna= false kind= icann
		    url_test.go:117: elements  tld= co.uk apex= www.co.uk host= one.www.co.uk port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  165.44.22.11 success= true
		    url_test.go:116: flags     compare= true ip= true idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host= 165.44.22.11 port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  165.44.22.99             success= true
		    url_test.go:116: flags     compare= true ip= true idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host= 165.44.22.99 port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  16.88.22.44:1234/path/page success= true
		    url_test.go:116: flags     compare= true ip= true idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host= 16.88.22.44 port= 1234 path= path/page
		    url_test.go:118:
		    url_test.go:115: origional  acdf::1212 success= true
		    url_test.go:116: flags     compare= true ip= true idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host= acdf::1212 port=  path=
		    url_test.go:118:
		    url_test.go:115: origional  [acca::2222]:5678/path/page success= true
		    url_test.go:116: flags     compare= true ip= true idna= false kind= bad
		    url_test.go:117: elements  tld=  apex=  host= acca::2222 port= 5678 path= path/page
		    url_test.go:118:
		--- PASS: TestURL (0.00s)
		PASS
	*/

	u := urlx.NewURL()
	t.Log("tld:", u.Len())

	// u.InvalidTLD() // allow invalid tld

	for _, v := range []string{
		"14x33",                                // bad; no tld
		"com",                                  // bad; icann tld only
		"api.c",                                // bad; not tld
		"duckdns.org",                          // bad; publicsuffix tld only
		"vfs.cloud9.eu-west-1.amazonaws.com",   // bad; publicsuffix tld only
		"exàmple.com",                          // idna
		"http://www.example.com",               // scheme+www+apex
		"blog.example.com:80/path/page",        // full
		"click.api.example.com:8888/path/page", // full
		"www.fr",                               // icann tld registrar edge case exception
		"www.co.uk",                            // icann tld registrar edge case exception
		"www.duckdns.org",                      // private tld edge case exception
		"one.www.co.uk",                        // not an icann edge case; apex is edge case exception
		"165.44.22.11",                         // ipv4
		"165.44.22.99\t\t",                     // ipv4 whitespace
		"16.88.22.44:1234/path/page",           // ipv4+full
		"acdf::1212",                           // ipv6
		"[acca::2222]:5678/path/page",          // ipv6+full
	} {

		t.Log("origional ", v, "success=", u.Parse(&v))
		t.Log("flags     compare=", u.Compare(), "ip=", u.IP, "idna=", u.IDNA, "kind=", urlx.Kind(u.Kind))
		t.Log("elements  tld=", u.TLD, "apex=", u.Apex, "host=", u.Host, "port=", u.Port, "path=", u.Path)
		t.Log()

	}

}

// TestParseIgnoresEmbeddedScheme covers a real production case (netstar-labs/clab,
// dispute-submission text with a link-protection wrapper or tracking-redirect
// path): the raw field is not a clean URL, and it carries a SECOND, later
// "://" from an embedded destination link. Parse must anchor on the real host
// at the start of the string and ignore that later scheme - not silently
// discard the real host and return the embedded one, which is what happened
// before this test existed: a dispute against "herchenbach.us" was reported
// against facebook.com's tranco rank instead, because the raw field ended in
// "...>[cid:...]<https://www.facebook.com/herchenbachgmbh".
func TestParseIgnoresEmbeddedScheme(t *testing.T) {

	u := urlx.NewURL()

	for _, tc := range []struct {
		name, in, wantApex string
	}{
		{
			name:     "html link-protection artifact",
			in:       "herchenbach.us/e3t/ctc/w1+23284/cmrlq04/jkm2>[cid:bc50f260]<https://www.facebook.com/herchenbachgmbh",
			wantApex: "herchenbach.us",
		},
		{
			name:     "redirect-service path embedding a destination scheme",
			in:       "jp.av4us.vip/v/s://www.dailymail.co.uk/news/article-7448291/",
			wantApex: "av4us.vip",
		},
		{
			name:     "genuine leading scheme still strips (regression guard)",
			in:       "http://www.example.com",
			wantApex: "example.com",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			in := tc.in
			if !u.Parse(&in) {
				t.Fatalf("Parse(%q) failed", tc.in)
			}
			if u.Apex != tc.wantApex {
				t.Errorf("Parse(%q).Apex = %q, want %q", tc.in, u.Apex, tc.wantApex)
			}
		})
	}
}

// TestParseAuthorityGrammar locks in WHATWG authority parsing, which is what a
// browser navigates by: backslash folds to '/', the authority ends at the first
// '/', '?' or '#', userinfo ends at the LAST '@', the host is percent-decoded
// before use, tab/CR/LF are deleted, IPv4 accepts every inet_aton radix form,
// and the special schemes tolerate missing or extra slashes. Each row where a
// naive byte scan disagrees is a host misattribution - an innocent domain
// blamed for an attacker's host, or an attacker host hidden where the parser
// reads a path. Expected hosts match github.com/netstar-labs/normie.Split.
func TestParseAuthorityGrammar(t *testing.T) {

	u := urlx.NewURL()

	for _, tc := range []struct {
		in, host, port, path string
		ok                   bool
	}{
		// authority terminators and userinfo
		{in: `http://example.com\@evil.com/`, host: "example.com", path: "@evil.com", ok: true},
		{in: `http://example.com#@evil.com`, host: "example.com", ok: true},
		{in: `http://example.com?@evil.com`, host: "example.com", ok: true},
		{in: `http://example.com?x=a@evil.com/p`, host: "example.com", ok: true},
		{in: `http://a@b@evil.com/`, host: "evil.com", ok: true},
		{in: `http://u:p@q@evil.com:8443/x`, host: "evil.com", port: "8443", path: "x", ok: true},
		{in: `http://example.com:80@evil.com/`, host: "evil.com", ok: true},
		{in: `https://example.com\x/y`, host: "example.com", path: "x/y", ok: true},

		// a "://" after the authority is content, never a scheme
		{in: `example.com?x=http://foo.com`, host: "example.com", ok: true},
		{in: `http://example.com?x=http://evil.com/`, host: "example.com", ok: true},

		// special schemes tolerate missing/extra slashes; protocol-relative
		{in: `http:/example.com/`, host: "example.com", ok: true},
		{in: `http:example.com/`, host: "example.com", ok: true},
		{in: `http:///example.com/`, host: "example.com", ok: true},
		{in: `//example.com/x`, host: "example.com", path: "x", ok: true},
		{in: `http://example.com:8080/x`, host: "example.com", port: "8080", path: "x", ok: true},

		// host byte cleaning: tab/CR/LF deleted, percent-escapes decoded
		{in: "http://exa\tmple.com/", host: "example.com", ok: true},
		{in: `http://%65xample.com/`, host: "example.com", ok: true},
		{in: `http://example.com%2fevil.com/`}, // an encoded delimiter decoded into a host is rejected

		// inet_aton IPv4 forms are the same address as the dotted quad
		{in: `http://134744072/`, host: "8.8.8.8", ok: true},
		{in: `http://0x08.8.8.8/`, host: "8.8.8.8", ok: true},
		{in: `http://010.8.8.8/`, host: "8.8.8.8", ok: true},
		{in: `http://8.8.2056/`, host: "8.8.8.8", ok: true},
		{in: `http://2130706433/`}, // loopback via dword: rejected like 127.0.0.1
		{in: `http://127.1/`},      // loopback via short form
		{in: `http://123/`},        // 0.0.0.0/8 is never a destination
		{in: `http://1.2.3.4.5/`},  // five parts: not an address, and no valid tld

		// ipv6: bracketed with and without a port (the unported form lost its last byte)
		{in: `http://[2606:4700::1111]/`, host: "2606:4700::1111", ok: true},
		{in: `[2606:4700::1111]`, host: "2606:4700::1111", ok: true},
		{in: `[2606:4700::1111]:8443/x`, host: "2606:4700::1111", port: "8443", path: "x", ok: true},
	} {
		t.Run(tc.in, func(t *testing.T) {
			in := tc.in
			ok := u.Parse(&in)
			if ok != tc.ok || u.Host != tc.host || u.Port != tc.port || u.Path != tc.path {
				t.Errorf("Parse(%q) = ok=%v host=%q port=%q path=%q; want ok=%v host=%q port=%q path=%q",
					tc.in, ok, u.Host, u.Port, u.Path, tc.ok, tc.host, tc.port, tc.path)
			}
		})
	}
}

// TestParseRejectsEmptyHost: a URL whose authority is empty has no host to
// report. With the tld check relaxed (InvalidTLD) Parse used to return true and
// hand back an empty Host for these, which a caller cannot tell from "no match".
func TestParseRejectsEmptyHost(t *testing.T) {

	u := urlx.NewURL().InvalidTLD()

	for _, in := range []string{"", "/", "/path", "http://", "http:///", "#frag", "?q", "@", "http://@/"} {
		t.Run(in, func(t *testing.T) {
			v := in
			if u.Parse(&v) {
				t.Errorf("Parse(%q) = true with host=%q; want false", in, u.Host)
			}
		})
	}
}

// TestParseIDNAProducedIPv4: UTS-46 mapping folds fullwidth digits and ideographic
// full stops to ASCII, so a host that is not an address before mapping can be one
// after it, and a browser maps before it parses an address. Probing only before
// mapping lets a loopback literal through as an ordinary host once the tld check
// is relaxed.
func TestParseIDNAProducedIPv4(t *testing.T) {

	strict, relaxed := urlx.NewURL(), urlx.NewURL().InvalidTLD()

	for _, tc := range []struct {
		in, host string
		ok       bool
	}{
		{in: "http://８.８.８.８/", host: "8.8.8.8", ok: true},
		{in: "http://８.８.２０５６/", host: "8.8.8.8", ok: true},
		{in: "http://１２７.０.０.１/"},  // loopback once mapped
		{in: "http://0x7f。0。0。1/"}, // loopback once mapped, radix form
		{in: "http://１０.０.０.１/"},   // private once mapped
	} {
		for name, u := range map[string]*urlx.URL{"strict": strict, "relaxed": relaxed} {
			t.Run(name+"/"+tc.in, func(t *testing.T) {
				in := tc.in
				ok := u.Parse(&in)
				if ok != tc.ok || u.Host != tc.host || ok && !u.IP {
					t.Errorf("Parse(%q) = ok=%v host=%q ip=%v; want ok=%v host=%q ip=%v", tc.in, ok, u.Host, u.IP, tc.ok, tc.host, tc.ok)
				}
			})
		}
	}
}

// TestParseIPv6Canonical: an address is one key however it is written. Uppercase,
// uncompressed and IPv4-mapped spellings collapse to the form a lookup (and the
// IPv4 spelling of the same host) agrees on, so a spelling cannot evade a list;
// a mapped loopback or private address is rejected like its IPv4 form.
func TestParseIPv6Canonical(t *testing.T) {

	u := urlx.NewURL()

	for _, tc := range []struct {
		in, host string
		ok       bool
	}{
		{in: "2606:4700:0:0:0:0:0:1111", host: "2606:4700::1111", ok: true},
		{in: "http://[2606:4700::ABCD]/", host: "2606:4700::abcd", ok: true},
		{in: "::ffff:8.8.8.8", host: "8.8.8.8", ok: true},
		{in: "http://[::ffff:808:808]/", host: "8.8.8.8", ok: true},
		{in: "http://[::FFFF:8.8.8.8]:80/", host: "8.8.8.8", ok: true},
		{in: "http://[::ffff:127.0.0.1]/"},
		{in: "http://[::ffff:10.0.0.1]/"},
		{in: "http://[0:0:0:0:0:0:0:1]/"},
	} {
		t.Run(tc.in, func(t *testing.T) {
			in := tc.in
			ok := u.Parse(&in)
			if ok != tc.ok || u.Host != tc.host || ok && !u.IP {
				t.Errorf("Parse(%q) = ok=%v host=%q ip=%v; want ok=%v host=%q", tc.in, ok, u.Host, u.IP, tc.ok, tc.host)
			}
		})
	}
}

// TestParseNonNavigableSchemes: only the web schemes are navigable URLs, so any
// other scheme must not have a host attributed to it. "javascript://example.com/
// %0aalert(1)" is a script, "mailto:a@example.com" an address. The forms that
// merely look like a scheme (host:port, host:junk, bare IPv6) still parse.
func TestParseNonNavigableSchemes(t *testing.T) {

	relaxed := urlx.NewURL().InvalidTLD() // so only the scheme rule can reject these

	for _, in := range []string{
		`javascript://example.com/%0aalert(1)`, `custom://example.com/x`, `ssh://example.com/`, `git+https://example.com/r`,
		`mailto:a@example.com`, `javascript:alert(1)`, `data:text/html,x.example.com`, `tel:+15551234567`,
		`data:a@b@example.com:80`, `user:pass@example.com`, `user:pass@example.com:8080/x`,
	} {
		t.Run("reject/"+in, func(t *testing.T) {
			v := in
			if relaxed.Parse(&v) {
				t.Errorf("Parse(%q) = true host=%q; want false", in, relaxed.Host)
			}
		})
	}

	u := urlx.NewURL()
	for _, tc := range []struct{ in, host, port string }{
		{`file://example.com/x`, "example.com", ""},
		{`ftp://example.com/x`, "example.com", ""},
		{`ws://example.com:8080/x`, "example.com", "8080"},
		{`example.com:8080/x`, "example.com", "8080"},
		{`example.com:80a`, "example.com", "80a"},
		{`abcd:ef01:2345::1`, "abcd:ef01:2345::1", ""},
	} {
		t.Run("accept/"+tc.in, func(t *testing.T) {
			v := tc.in
			if !u.Parse(&v) || u.Host != tc.host || u.Port != tc.port {
				t.Errorf("Parse(%q) = host=%q port=%q; want host=%q port=%q", tc.in, u.Host, u.Port, tc.host, tc.port)
			}
		})
	}
}

// TestParseCollapseDots: empty labels are collapsed, not carried into the host or
// rejected, as the GSB canonicalization does: leading dots dropped, runs folded to
// one, trailing dots dropped. "evil..com" is the host a lookup keys as "evil.com";
// keeping the empty label means it never matches.
func TestParseCollapseDots(t *testing.T) {

	u := urlx.NewURL()

	for _, tc := range []struct {
		in, host string
		ok       bool
	}{
		{in: "http://.example.com/", host: "example.com", ok: true},
		{in: "http://..example.com/", host: "example.com", ok: true},
		{in: "http://example..com/", host: "example.com", ok: true},
		{in: "http://a...b....example.com./x", host: "a.b.example.com", ok: true},
		{in: "http://ex%2e%2eample.com/", host: "ex.ample.com", ok: true},
		{in: "http://example。。com/", host: "example.com", ok: true},
		{in: "a．．b.com", host: "a.b.com", ok: true},
		{in: "http://./", host: ""},
		{in: "http://.../", host: ""},
	} {
		t.Run(tc.in, func(t *testing.T) {
			in := tc.in
			ok := u.Parse(&in)
			if ok != tc.ok || u.Host != tc.host {
				t.Errorf("Parse(%q) = ok=%v host=%q; want ok=%v host=%q", tc.in, ok, u.Host, tc.ok, tc.host)
			}
		})
	}
}

// TestParseRejectsForbiddenHostBytes: a hostname containing a space, control or URL
// delimiter is not a host, and a browser does not navigate to one. Parse used to
// accept "exa mple.com" or "exa<mple.com" whole (only the ends were trimmed), so a
// dirty fragment was reported as a host instead of being rejected.
func TestParseRejectsForbiddenHostBytes(t *testing.T) {

	for name, u := range map[string]*urlx.URL{
		"strict": urlx.NewURL(), "relaxed": urlx.NewURL().InvalidTLD(), "noidna": urlx.NewURL().InvalidTLD().NoIDNA(),
	} {
		for _, in := range []string{
			"http://exa mple.com/", "http://exa<mple.com/", `http://exa"mple.com/`, "http://exa|mple.com/",
			"http://exa^mple.com/", "http://exa>mple.com/", "data:134744072:65535", "data:example.com:",
		} {
			t.Run(name+"/"+in, func(t *testing.T) {
				v := in
				if u.Parse(&v) {
					t.Errorf("Parse(%q) = true host=%q; want false", in, u.Host)
				}
			})
		}
	}
}
