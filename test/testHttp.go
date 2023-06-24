package main

import (
	"encoding/json"
	"fmt"
)

type rt struct {
	name string
}

func main() {
	marshal, err := json.Marshal(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	var resultData rt
	err = json.Unmarshal(marshal, &resultData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultData)
}
