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
