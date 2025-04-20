
# urlx

The urlx package is a simple url standardizer that parses a url into parts and will set several flags.
 url.IP or url.IDNA flag when detected. 

Parse generally confirms the hostname has the required format and generated the apex, host, and tld with the tld kind for domain and reports if an IPv4/6 address.


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


