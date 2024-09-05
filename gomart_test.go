package gomart

import "testing"

func TestInit(t *testing.T) {
	token := "token"
	client := Init(token)
	t.Logf("%+v", client)
}
