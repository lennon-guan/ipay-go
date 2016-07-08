package ipay

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func getAppId() string {
	appid := os.Getenv("APPID")
	if appid == "" {
		panic("please set env var APPID")
	}
	return appid
}

func TestCreateOrder(t *testing.T) {
	h, err := NewIpayHelperWithPem(getAppId(), "test_private_key.pem", "test_public_key.pem")
	if err != nil {
		t.Error(err)
	}
	transId, err := h.CreateIpayOrder(1, "打架=服务费", "A34544", 3.45, "10086", "123", "")
	if err != nil {
		t.Error(err)
		return
	}
	payUrl, err := h.GetHtml5RedirectUrl(transId, "http://here", "")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("transid = %s\npayUrl = %s\n", transId, payUrl)
}

func prettyPrint(d interface{}) {
	jb, _ := json.MarshalIndent(d, "", "  ")
	fmt.Println(string(jb))
}

func TestQueryResult(t *testing.T) {
	h, err := NewIpayHelperWithPem(getAppId(), "test_private_key.pem", "test_public_key.pem")
	if err != nil {
		t.Error(err)
	}
	r, err := h.QueryResult("A34544")
	if err != nil {
		t.Error(err)
	}
	prettyPrint(r)
}

func TestParseNotifyInfo(t *testing.T) {

}
