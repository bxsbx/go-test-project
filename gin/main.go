package main

import (
	"StandardProject/gin/global"
	"StandardProject/gin/router"
	"fmt"
	"net/http"
	"time"
)

//go:generate go mod tidy

func main() {

	global.Init()

	routers := router.Routers()
	server := &http.Server{
		Addr:           "",
		Handler:        routers,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
