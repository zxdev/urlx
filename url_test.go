package urlx_test

import (
	"testing"

	"github.com/zxdev/urlx"
)

func TestURL(t *testing.T) {

	/*
		=== RUN   TestURL
		    url_test.go:102: origional  14x33 success= false
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host=  port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  com success= false
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host=  port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  api.c success= false
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host=  port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  duckdns.org success= false
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host=  port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  exàmple.com success= true
		    url_test.go:103: flags     compare= true ip= false idna= true kind= 1
		    url_test.go:104: elements  tld= com apex= xn--exmple-jta.com host= xn--exmple-jta.com port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  http://www.example.com success= true
		    url_test.go:103: flags     compare= true ip= false idna= false kind= 1
		    url_test.go:104: elements  tld= com apex= example.com host= example.com port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  blog.example.com:80/path/page success= true
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 1
		    url_test.go:104: elements  tld= com apex= example.com host= blog.example.com port= 80 path= path/page
		    url_test.go:105:
		    url_test.go:102: origional  click.api.example.com:8888/path/page success= true
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 1
		    url_test.go:104: elements  tld= com apex= example.com host= click.api.example.com port= 8888 path= path/page
		    url_test.go:105:
		    url_test.go:102: origional  www.fr success= true
		    url_test.go:103: flags     compare= true ip= false idna= false kind= 1
		    url_test.go:104: elements  tld= fr apex= www.fr host= www.fr port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  www.co.uk success= true
		    url_test.go:103: flags     compare= true ip= false idna= false kind= 1
		    url_test.go:104: elements  tld= co.uk apex= www.co.uk host= www.co.uk port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  www.duckdns.org success= true
		    url_test.go:103: flags     compare= true ip= false idna= false kind= 2
		    url_test.go:104: elements  tld= duckdns.org apex= www.duckdns.org host= www.duckdns.org port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  one.www.co.uk success= true
		    url_test.go:103: flags     compare= false ip= false idna= false kind= 1
		    url_test.go:104: elements  tld= co.uk apex= www.co.uk host= one.www.co.uk port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  165.44.22.11 success= true
		    url_test.go:103: flags     compare= true ip= true idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host= 165.44.22.11 port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  16.88.22.44:1234/path/page success= true
		    url_test.go:103: flags     compare= true ip= true idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host= 16.88.22.44 port= 1234 path= path/page
		    url_test.go:105:
		    url_test.go:102: origional  acdf::1212 success= true
		    url_test.go:103: flags     compare= true ip= true idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host= acdf::1212 port=  path=
		    url_test.go:105:
		    url_test.go:102: origional  [acca::2222]:5678/path/page success= true
		    url_test.go:103: flags     compare= true ip= true idna= false kind= 0
		    url_test.go:104: elements  tld=  apex=  host= acca::2222 port= 5678 path= path/page
		    url_test.go:105:
		--- PASS: TestURL (0.00s)
	*/

	u := urlx.Parser(nil)
	//u.ValidTLD() // allow invalid tld

	for _, v := range []string{
		"14x33",                                // bad; no tld
		"com",                                  // bad; icann tld only
		"api.c",                                // bad; not tld
		"duckdns.org",                          // bad; publicsuffix tld only
		"exàmple.com",                          // idna
		"http://www.example.com",               // scheme+www+apex
		"blog.example.com:80/path/page",        // full
		"click.api.example.com:8888/path/page", // full
		"www.fr",                               // icann tld registrar edge case exception
		"www.co.uk",                            // icann tld registrar edge case exception
		"www.duckdns.org",                      // private tld edge case exception
		"one.www.co.uk",                        // not an icann edge case; apex is edge case exception
		"165.44.22.11",                         // ipv4
		"16.88.22.44:1234/path/page",           // ipv4+full
		"acdf::1212",                           // ipv6
		"[acca::2222]:5678/path/page",          // ipv6+full
	} {

		t.Log("origional ", v, "success=", u.Parse(&v))
		t.Log("flags     compare=", u.Compare(), "ip=", u.IP, "idna=", u.IDNA, "kind=", u.Kind)
		t.Log("elements  tld=", u.TLD, "apex=", u.Apex, "host=", u.Host, "port=", u.Port, "path=", u.Path)
		t.Log()

	}

}
