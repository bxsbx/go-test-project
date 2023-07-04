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

var client map[string]*http.Client

//func init() {
//	client = make(map[string]*http.Client)
//	client[DEFAULT] = newClient(1)
//	client[HOMEWORK] = newClient(1)
//}

func DefaultClient() *http.Client {
	return client[DEFAULT]
}

func GetClient(clientKey string) *http.Client {
	return client[clientKey]
}
