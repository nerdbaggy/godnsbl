/*
Gets the results and reports them in json
  package main
  import (
      "fmt"
      "encoding/json"
      "github.com/nerdbaggy/godnsbl"
  )
  func main() {
    resp := godnsbl.CheckBlacklist("127.0.0.2")
    jsonResponse, _ := json.Marshal(resp)
    fmt.Printf("%s", jsonResponse)
  }
Output:
{
"Err": "",
"Listed": true,
"Count": 32,
"Total": 37,
"Responses": [{
  "Status": "ok",
  "Msg": "",
  "Listed": true,
  "Name": "bl.spamcop.net",
  "RespTime": 56,
  "Resp": "127.0.0.2"
}, {
  "Status": "ok",
  "Msg": "",
  "Listed": true,
  "Name": "dnsrbl.org",
  "RespTime": 55,
  "Resp": "127.0.0.2"
}, {
  "Status": "ok",
  "Msg": "",
  "Listed": true,
  "Name": "misc.dnsbl.sorbs.net",
  "RespTime": 55,
  "Resp": "127.0.0.4"
}, {
  "Status": "ok",
  "Msg": "",
  "Listed": true,
  "Name": "smtp.dnsbl.sorbs.net",
  "RespTime": 56,
  "Resp": "127.0.0.5"
}, {
  "Status": "ok",
  "Msg": "",
  "Listed": true,
  "Name": "escalations.dnsbl.sorbs.net",
  "RespTime": 57,
  "Resp": "127.0.0.6"
}, {
  "Status": "ok",
  "Msg": "",
  "Listed": true,
  "Name": "dnsbl-2.uceprotect.net",
  "RespTime": 57,
  "Resp": "127.0.0.2"
}]
}
*/
package godnsbl
