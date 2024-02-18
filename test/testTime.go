package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type JiBu struct {
	A  int
	Ok *time.Time `json:"ok"`
}

func main() {
	str := "{ \"A\":98,\"ok\":\"2022-10-15T20:00:00Z\"}"
	item := JiBu{}
	decoder := json.NewDecoder(bytes.NewReader([]byte(str)))
	//err := json.Unmarshal([]byte(str), &item)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if err := decoder.Decode(&item); err != nil {
		log.Fatal(err)
	}
	fmt.Println(item)

	marshal, _ := json.Marshal(item)
	fmt.Println(string(marshal))
}
