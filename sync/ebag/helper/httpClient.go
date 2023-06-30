package helper

import (
	"crypto/tls"
	"net/http"
	"time"
)

func GetClient(timeout int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			DisableCompression:    false,
			ResponseHeaderTimeout: time.Second * time.Duration(timeout),
		},
	}
}
