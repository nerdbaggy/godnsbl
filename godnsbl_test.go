package godnsbl

import (
  "testing"
  "sync"
)

func TestGetDnsblRespFunc(t *testing.T) {
  var wg sync.WaitGroup
  var resp DnsblReturn
  wg.Add(1)
  go getDnsblResp(&wg, &resp, getFlipIP("127.0.0.2"), "b.barracudacentral.org")
  wg.Wait()

  if len(resp.Responses) != 1{
    t.Errorf("Expected length of responses to be 1, got: %d", len(resp.Responses))
  }

  if resp.Responses[0].Status != "ok"{
    t.Errorf("Expected status to be ok, got : %s", resp.Responses[0].Status)
  }

  if resp.Responses[0].Msg != ""{
    t.Errorf("No message expected, got : %s", resp.Responses[0].Msg)
  }

  if resp.Responses[0].Listed != true{
    t.Errorf("Expected blacklist to be true, got : %s", resp.Responses[0].Listed)
  }

  if resp.Responses[0].Name != "b.barracudacentral.org"{
    t.Errorf("Expected name to be b.barracudacentral.org, got : %s", resp.Responses[0].Name)
  }

  if resp.Responses[0].RespTime == 0{
    t.Errorf("Response time should not be zero")
  }

  if resp.Responses[0].Resp != "127.0.0.2"{
    t.Errorf("Resp should be 127.0.0.2, got: %s", resp.Responses[0].Resp)
  }

}


func TestCheckBlacklistFunc(t *testing.T) {
  resp := CheckBlacklist("127.0.0.2")

  if len(resp.Responses) != len(BlacklistDomains){
    t.Errorf("Expected length of responses to be %d, got: %d", len(BlacklistDomains), len(resp.Responses))
  }

  if resp.Listed != true{
    t.Errorf("Expected blacklist to be true, got : %s", resp.Listed)
  }

}

func TestGetflipIPFunc(t *testing.T) {
  flipped := getFlipIP("10.11.12.13")
  if flipped != "13.12.11.10"{
		t.Errorf("IP did not flip properly, expecting 13.12.11.10 got: %s", flipped)
	}
}
