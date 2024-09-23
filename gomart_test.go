package gomart

import (
	"os"
	"testing"
)

var token string

func TestMain(m *testing.M) {
	token = os.Getenv("MARKET_TOKEN")
	m.Run()
}

func TestToken(t *testing.T) {
	client, err := Init(token)
	if err != nil {
		t.Error(err)
	}
	for _, campaign := range client.campaigns {
		t.Logf("%+v", campaign.Business.Id)
	}
}
