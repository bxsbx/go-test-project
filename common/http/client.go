package http

import (
	"net/http"
	"time"
)

const (
	DEFAULT  = "default"
	HOMEWORK = "homework"
	LOGIN    = "login"
)

// 秒单位
func newClient(timeOut int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * time.Duration(timeOut),
		},
	}
}

var client map[string]*Client

func init() {
	client = make(map[string]*Client)
	client[DEFAULT] = &Client{RespType: 1, Client: newClient(3)}
	client[HOMEWORK] = &Client{RespType: 2, Client: newClient(1)}
}

func DefaultClient() *Client {
	return client[DEFAULT]
}

func GetClient(clientKey string) *Client {
	return client[clientKey]
}
