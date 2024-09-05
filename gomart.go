package gomart

import "net/http"

type Client struct {
	client http.Client
	token  string
}

func Init(token string) *Client {
	var client = http.Client{}
	return &Client{
		client,
		token,
	}
}
