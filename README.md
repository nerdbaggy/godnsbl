# godnsbl [![Build Status](https://travis-ci.org/nerdbaggy/godnsbl.svg?branch=master)](https://travis-ci.org/nerdbaggy/godnsbl) [![GoDoc](https://godoc.org/github.com/nerdbaggy/godnsbl?status.svg)](https://godoc.org/github.com/nerdbaggy/godnsbl)

### [Full Documentation](https://godoc.org/github.com/nerdbaggy/godnsbl#DnsblReturn)

godnsbl will take an IP and tell you which dnsbl(dnsrbl) that IP is listed on. 

## Output
Data is returned as a struct. [Documentation](https://godoc.org/github.com/nerdbaggy/godnsbl#DnsblReturn)
```
type DnsblReturn struct {
    Err       string      // Will always be "" unless there is an error
    Listed    bool        // If any of the dnsrbl lists contain this IP
    Count     int         // How many have blacklisted this IP
    Total     int         // How many total dnsrbl were checked
    Responses []DnsblData // List of all the responses
}
```

If the output was rendered to json it would look like the following below. This example would normally list 37 examples but I truncated it to save space to 3.
```
{
	"Err": "",
	"Listed": true,
	"Count": 32,
	"Total": 37,
	"Responses": [{
		"Status": "ok",
		"Msg": "",
		"Listed": true,
		"Name": "http.dnsbl.sorbs.net",
		"RespTime": 43,
		"Resp": "127.0.0.2"
	}, {
		"Status": "ok",
		"Msg": "",
		"Listed": true,
		"Name": "web.dnsbl.sorbs.net",
		"RespTime": 41,
		"Resp": "127.0.0.7"
	}, {
		"Status": "ok",
		"Msg": "",
		"Listed": true,
		"Name": "bl.spamcop.net",
		"RespTime": 52,
		"Resp": "127.0.0.2"
	}]
}
```
## Currently Checks the following DNSBL
- b.barracudacentral.org
- bl.spamcop.net
- dnsrbl.org
- babl.rbl.webiron.net
- cabl.rbl.webiron.net
- stabl.rbl.webiron.net
- crawler.rbl.webiron.net
- bad.psky.me
- http.dnsbl.sorbs.net
- socks.dnsbl.sorbs.net
- misc.dnsbl.sorbs.net
- smtp.dnsbl.sorbs.net
- web.dnsbl.sorbs.net
- spam.dnsbl.sorbs.net
- escalations.dnsbl.sorbs.net
- zombie.dnsbl.sorbs.net
- recent.spam.dnsbl.sorbs.net
- dyna.spamrats.com
- spam.abuse.ch
- rbl.interserver.net
- tor.dan.me.uk
- torexit.dan.me.uk
- dnsbl-1.uceprotect.net
- dnsbl-2.uceprotect.net
- dnsbl-3.uceprotect.net
- cbl.abuseat.org
- spam.spamrats.com
- ips.backscatterer.org
- truncate.gbudb.net
- psbl.surriel.com
- db.wpbl.info
- bl.spamcannibal.org
- dnsbl.inps.de
- bl.blocklist.de
- rbl.megarbl.net
- all.s5h.net
- srnblack.surgate.net
