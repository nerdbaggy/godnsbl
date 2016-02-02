//Package godnsbl provides a way to query if an IP is on any dnsbl blacklist
package godnsbl

import (
  "sync"
  "strings"
  "fmt"
  "net"
  "time"
)

var blacklistDomains = []string{
  "b.barracudacentral.org",
  "bl.spamcop.net",
  "dnsrbl.org",
  "babl.rbl.webiron.net",
  "cabl.rbl.webiron.net",
  "stabl.rbl.webiron.net",
  "crawler.rbl.webiron.net",
  "bad.psky.me",
  "http.dnsbl.sorbs.net",
  "socks.dnsbl.sorbs.net",
  "misc.dnsbl.sorbs.net",
  "smtp.dnsbl.sorbs.net",
  "web.dnsbl.sorbs.net",
  "spam.dnsbl.sorbs.net",
  "escalations.dnsbl.sorbs.net",
  "zombie.dnsbl.sorbs.net",
  "recent.spam.dnsbl.sorbs.net",
  "dyna.spamrats.com",
  "spam.abuse.ch",
  "rbl.interserver.net",
  "tor.dan.me.uk",
  "torexit.dan.me.uk",
  "dnsbl-1.uceprotect.net",
  "dnsbl-2.uceprotect.net",
  "dnsbl-3.uceprotect.net",
  "cbl.abuseat.org",
  "spam.spamrats.com",
  "ips.backscatterer.org",
  "truncate.gbudb.net",
  "psbl.surriel.com",
  "db.wpbl.info",
  "bl.spamcannibal.org",
  "dnsbl.inps.de",
  "bl.blocklist.de",
  "rbl.megarbl.net",
  "all.s5h.net",
  "srnblack.surgate.net",
}

// A dnsblData contains all the data about a particular dnsbl list
type dnsblData struct {
  Status string
  Msg string
  Listed bool
  Name string
  RespTime int64
  Resp string
}

// dnsblReturn contains a list of all the checked lists and if its on any lists
type dnsblReturn struct{
  Listed bool
  Count int
  Total int
  Responses []dnsblData
}


func getDnsblResp(wg *sync.WaitGroup, resp *dnsblReturn, ip string, domain string) {
  // Sweet timing function to see how long it takes to get a response
  st := time.Now()

  // Get the host for the format <reverse ip>.domain
  ips, err := net.LookupIP(fmt.Sprintf("%s.%s", ip, domain))

  // See how long it took the request
  et := time.Since(st).Nanoseconds()/1000000

  // Default values to insert
  listed := false
  respIPs := ""
  respStat := "ok"
  respErrMsg := ""

  // If it gets an IP back, it is listed in the database. Only get the first IP per RFC
  if len(ips) > 0 {
    listed = true
    respIPs = ips[0].String()
  }

  // Only show the error if it isn't the generic no domain when the listing isnt listed
  if err != nil && err.Error() != fmt.Sprintf("lookup %s.%s: no such host", ip, domain) && err.Error() != fmt.Sprintf("lookup %s.%s: No address associated with hostname", ip, domain){
    respStat = "error"
    respErrMsg = err.Error()
  }

  // Could of used channels but this works for now
  resp.Responses = append(resp.Responses, dnsblData{respStat, respErrMsg, listed, domain, et, respIPs})

  // Bows
  wg.Done()
}

func CheckBlacklist(ip string) (dnsblReturn){
  var wg sync.WaitGroup
  wg.Add(len(blacklistDomains))
  var resp dnsblReturn

  // Flip the IP
  flipIP := getFlipIP(ip)

  // Makes a go routine for every domain and does the check
  for _, dom := range blacklistDomains {
    go getDnsblResp(&wg, &resp, flipIP, dom)
  }

  // They do take a while to return all
  wg.Wait()

  // Loop through each reply and see if the IP is on any blacklists
  for _, check := range resp.Responses{
    if check.Listed == true{
      resp.Listed = true
      resp.Count += 1
    }
  }
  resp.Total = len(blacklistDomains)

  // We are finished!
  return resp
}

func getFlipIP(ip string) (string){
  // Splits the IP by the period
  brokenIP := strings.Split(ip, ".")
  // Flips the IP to be in proper RFC format
  return fmt.Sprintf("%s.%s.%s.%s", brokenIP[3], brokenIP[2], brokenIP[1], brokenIP[0])
}
