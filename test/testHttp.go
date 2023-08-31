package main

import (
	"StandardProject/common/errorz"
	myHttp "StandardProject/common/http"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type T2 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

var url = "http://127.0.0.1:8888/v1/student/web/test/test1"

func Request(j int) (T2, error) {
	queryMap := make(map[string]string)
	queryMap["id"] = strconv.Itoa(j)
	var r T2

	client := &http.Client{
		Transport: &http.Transport{
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * time.Duration(3),
		},
	}
	err := myHttp.Get(url, queryMap, &r, 0, client, context.Background())
	//err := myHttp.Get(url, queryMap, &r, 0, myHttp.DefaultClient(), context.Background())
	//fmt.Println(r)
	return r, err
}

func main() {
	//http.DefaultClient()

	//Request(3)
	errNum, successNum := 0, 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	tempMap := make(map[string]string)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			r, err := Request(j)
			if err != nil {
				if myErr, ok := err.(*errorz.MyError); ok {
					//if j%2 == 0 {
					fmt.Println(j, myErr.Unwrap())
					//}
				}
				mu.Lock()
				errNum++
				mu.Unlock()
			} else {
				time.Sleep(2 * time.Second)
				mu.Lock()
				tempMap[r.Data] = r.Msg
				successNum++
				mu.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(len(tempMap))
	fmt.Println(errNum, successNum)
}
