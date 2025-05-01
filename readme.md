
# urlx

The urlx package is a simple url standardizer that parses a url into its constituant parts and will set several flags as part of the parser such as url.IP or url.IDNA flag when detected. The parser can be modified with various toggles, such as noWWW, noPort, noPage as well as various controls flags such as onlyIP, onlyApex, or onlyHost.

The Parser is used to generally confirms the hostname has the required format and generates the tld, apex, and host along with the tld kind for the domain and reports only the host for an IPv4/6 address.

The package provides a public suffix aware url parser based on data from https://publicsuffix.org/ that can be merged with a customlist.

A public suffix is one under which Internet users can directly register names. It is related to, but different from, a TLD (top level domain); "com" is a TLD (top level domain) which means it has no dots and "com" is also a public suffix. "au" is another TLD, again because it has no dots, and while co.uk is not technically a TLD, is is a recognized TLD because that's the branching point for domain name registrars.

All of these domains have the same eTLD+1 or Apex hostname "amazon":
- "www.books.amazon.co.uk"
- "books.amazon.co.uk"
- "amazon.co.uk"

Specifically, the eTLD+1 is "amazon.co.uk", because the eTLD is "co.uk".

There is no closed form algorithm to calculate the eTLD of a domain. Instead, the calculation is data driven and this package uses Mozilla's PSL (Public Suffix List) to detect the tld and apex form when parsing the url. The list automatically refreshes every 72 hours according the the foundations recommendation, and an additional custom list can be created and utilized as well to handle other unique situations.

The Compare method is a boolean equivalence comparison test for Apex and Host domain names. It's best to test for an IP first and take action at that branch point, however testing with an IP will report eqivalence even though the apex field is empty becuase the intended logic is that when the host and apex are different a different or secondary action would be taken based upon this testing so plan the resultant logic flow accordingly.

```golang

// example
var u = urlx.Parser(nil)
url := "api.example.com/path/logo.jpg"
if u.Parse(&url) {
    if u.IP {
        fmt.Println("host is an IP address")
    }
    if u.Compare() {
        fmt.Println("host is the same as apex")
    }
    fmt.Println("apex=", u.Apex, "host=", u.Host)
}

// === RUN   TestOne
// apex= example.com host= api.example.com
// --- PASS: TestOne (0.00s)
// PASS

```


