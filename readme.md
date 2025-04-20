
# urlx

The urlx package is a simple url standardizer that parses a url into its constituant  parts and will set several flags as part of the parser such as url.IP or url.IDNA flag when detected. The parser can be modified with various toggles, such as noWWW, noPort, noPage as well as various controls flags such as onlyIP, onlyApex, or onlyHost.

The Parser is used to generally confirms the hostname has the required format and generates the tld, apex, and host along with the tld kind for the domain and reports only the host for an IPv4/6 address.

```golang

// example
var u = urlx.Parser(nil)
url := "api.example.com/path/logo.jpg"
if u.Parse(&url) {
    if u.HostIsApex() {
        fmt.Println("host is the same as apex")
    }
    fmt.Println("apex=", u.Apex, "host=", u.Host)
}

// === RUN   TestOne
// apex= example.com host= api.example.com
// --- PASS: TestOne (0.00s)
// PASS

```


