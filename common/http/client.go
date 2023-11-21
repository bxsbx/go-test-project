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

var client map[string]*Client
var mu sync.RWMutex

func init() {
	client = make(map[string]*Client)
	client[DEFAULT] = &Client{RespType: 0, Client: newClient(3)}
	client[HOMEWORK] = &Client{RespType: 2, Client: newClient(1)}
}

func DefaultClient() *Client {
	return client[DEFAULT]
}

func GetClient(clientKey string) *Client {
	if v, ok := client[clientKey]; ok {
		return v
	}
	return client[DEFAULT]
}

func GetOrCreateClientByPath(path string, timeOut int) *Client {
	mu.RLock()
	v, ok := client[path]
	mu.RUnlock()
	if ok {
		return v
	}
	mu.Lock()
	defer mu.Unlock()
	client_new := &Client{RespType: 0, Client: newClient(timeOut)}
	client[path] = client_new
	return client_new
}
