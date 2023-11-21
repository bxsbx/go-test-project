package http

import (
	"net/http"
	"sync"
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

var client struct {
	sync.RWMutex
	clientMap map[string]*Client
}

func init() {
	client.clientMap = make(map[string]*Client)
	client.clientMap[DEFAULT] = &Client{RespType: 0, Client: newClient(3)}
	client.clientMap[HOMEWORK] = &Client{RespType: 2, Client: newClient(1)}
}

func DefaultClient() *Client {
	return client.clientMap[DEFAULT]
}

func GetClient(clientKey string) *Client {
	if v, ok := client.clientMap[clientKey]; ok {
		return v
	}
	return client.clientMap[DEFAULT]
}

func GetOrCreateClientByPath(path string, timeOut int) *Client {
	client.RLock()
	v, ok := client.clientMap[path]
	client.RUnlock()
	if ok {
		return v
	}
	client.Lock()
	defer client.Unlock()
	client_new := &Client{RespType: 0, Client: newClient(timeOut)}
	client.clientMap[path] = client_new
	return client_new
}
