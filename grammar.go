package urlx

import (
	"net/netip"
	"strconv"
	"strings"
)

// schemeEnd returns the index of the ':' ending a leading RFC 3986 scheme
// (ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )), or -1 when s does not start with one.
func schemeEnd(s string) int {
	if s == "" || !isAlpha(s[0]) {
		return -1
	}
	for i := 1; i < len(s); i++ {
		switch c := s[i]; {
		case isAlpha(c), c >= '0' && c <= '9', c == '+', c == '-', c == '.':
		case c == ':':
			return i
		default:
			return -1
		}
	}
	return -1
}

func isAlpha(c byte) bool { return c|0x20 >= 'a' && c|0x20 <= 'z' }

// isSpecialScheme reports whether s is a WHATWG special scheme that always has an
// authority, other than file (whose host is usually empty).
func isSpecialScheme(s string) bool {
	switch strings.ToLower(s) {
	case "http", "https", "ws", "wss", "ftp":
		return true
	}
	return false
}

// trailingAlpha returns the run of ASCII letters at the end of s ("\"ssh" -> "ssh").
func trailingAlpha(s string) string {
	i := len(s)
	for i > 0 && isAlpha(s[i-1]) {
		i--
	}
	return s[i:]
}

// opaqueScheme reports whether s begins with a non-special scheme that has no
// authority: "mailto:a@b.com", "javascript:alert(1)", "data:text/html,...",
// "tel:+1555". Two shapes only look like "scheme:rest" and are NOT opaque: a
// "host.tld:port" or "host:junk" (a '.' in the candidate, or digits after the
// colon) and a bare IPv6 literal (several ':' and no '@'). A second ':' alongside
// an '@' is no exemption: "data:a@b@example.com:80" is not a visit to example.com,
// and schemeless "user:pass@host" is opaque, as in github.com/netstar-labs/normie.
func opaqueScheme(s string) bool {
	i := schemeEnd(s)
	if i <= 0 {
		return false
	}
	scheme, rest := s[:i], s[i+1:]
	if strings.IndexByte(scheme, '.') >= 0 || isSpecialScheme(scheme) || strings.EqualFold(scheme, "file") ||
		strings.HasPrefix(rest, "//") || portFollows(rest) {
		return false
	}
	end := strings.IndexAny(s, "/?#")
	if end < 0 {
		end = len(s)
	}
	auth := s[:end]
	return strings.Count(auth, ":") == 1 || strings.IndexByte(auth, '@') >= 0
}

// portFollows reports whether s begins with digits ended by a delimiter or the
// end of s, i.e. the preceding ':' separated a port, not a scheme.
func portFollows(s string) bool {
	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	return i > 0 && (i == len(s) || strings.IndexByte("/?#", s[i]) >= 0)
}

// collapseDots drops leading and trailing dots and folds each run of dots to one,
// as the GSB canonicalization does: "..a...b." is "a.b". Returns s itself when
// nothing changes.
func collapseDots(s string) string {
	if !strings.Contains(s, "..") && !strings.HasPrefix(s, ".") && !strings.HasSuffix(s, ".") {
		return s
	}
	b := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '.' && (len(b) == 0 || b[len(b)-1] == '.') {
			continue
		}
		b = append(b, s[i])
	}
	return strings.TrimRight(string(b), ".")
}

// forbiddenHostByte reports whether b is a WHATWG forbidden host code point: a
// control, space, or an authority/URL delimiter.
func forbiddenHostByte(b byte) bool {
	return b <= 0x20 || b == 0x7f || strings.IndexByte("\"#%/:<>?@[\\]^|", b) >= 0
}

// percentDecodeHost decodes %XX escapes in host. ok is false for a malformed
// escape or one that decodes to a forbidden host byte; a host with no '%' is
// returned unchanged.
func percentDecodeHost(host string) (string, bool) {
	if !strings.ContainsRune(host, '%') {
		return host, true
	}
	var b strings.Builder
	b.Grow(len(host))
	for i := 0; i < len(host); i++ {
		if host[i] != '%' {
			b.WriteByte(host[i])
			continue
		}
		if i+2 >= len(host) {
			return "", false // dangling '%'
		}
		n, err := strconv.ParseUint(host[i+1:i+3], 16, 8)
		if err != nil || forbiddenHostByte(byte(n)) {
			return "", false // malformed %XX, or an encoded delimiter/control
		}
		b.WriteByte(byte(n))
		i += 2
	}
	return b.String(), true
}

// parseIPv4Radix implements inet_aton semantics: one to four dot-separated
// parts, each independently radix-detected (0x hex, leading-0 octal, else
// decimal), with the final part absorbing all remaining octets.
//
// It mirrors github.com/netstar-labs/normie.ParseIPv4; keep the two in step.
func parseIPv4Radix(s string) (netip.Addr, bool) {
	// Every part of an inet_aton address starts with a digit, so any other lead
	// byte is a hostname. This runs on every host, so it must not allocate.
	if s == "" || len(s) > 63 || s[0] < '0' || s[0] > '9' {
		return netip.Addr{}, false
	}
	var vals [4]uint64
	n := 0
	for {
		if n == len(vals) {
			return netip.Addr{}, false // more than four parts
		}
		part, rest, more := strings.Cut(s, ".")
		v, ok := parseIPv4Part(part)
		if !ok {
			return netip.Addr{}, false
		}
		vals[n] = v
		n++
		if !more {
			break
		}
		s = rest
	}
	for i := 0; i < n-1; i++ { // every part but the last is a single octet
		if vals[i] > 0xff {
			return netip.Addr{}, false
		}
	}
	if rem := uint(4 - (n - 1)); vals[n-1] >= uint64(1)<<(8*rem) {
		return netip.Addr{}, false // the last part must fit the octets left
	}
	var word uint32
	for i := 0; i < n-1; i++ {
		word |= uint32(vals[i]) << (8 * (3 - uint(i)))
	}
	word |= uint32(vals[n-1])
	return netip.AddrFrom4([4]byte{byte(word >> 24), byte(word >> 16), byte(word >> 8), byte(word)}), true
}

// parseIPv4Part decodes one component with C radix rules. An empty component is
// invalid, which rejects "1..2". 12 bytes is the longest meaningful part (a full
// 32-bit value in octal, "037777777777"), so v cannot overflow uint64. Parsed by
// hand: strconv.ParseUint heap-allocates its error on every non-numeric label.
func parseIPv4Part(p string) (uint64, bool) {
	if p == "" || len(p) > 12 {
		return 0, false
	}
	base := uint64(10)
	switch {
	case len(p) > 2 && p[0] == '0' && p[1]|0x20 == 'x':
		base, p = 16, p[2:]
	case len(p) > 1 && p[0] == '0':
		base, p = 8, p[1:]
	}
	var v uint64
	for i := 0; i < len(p); i++ {
		c := p[i]
		var d uint64
		switch {
		case c >= '0' && c <= '9':
			d = uint64(c - '0')
		case c|0x20 >= 'a' && c|0x20 <= 'f':
			d = uint64(c|0x20-'a') + 10
		default:
			return 0, false
		}
		if d >= base {
			return 0, false
		}
		v = v*base + d
	}
	return v, true
}
