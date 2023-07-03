package main

import (
	"StandardProject/sync/ebag/questions"
	"fmt"
	"log"
)

func main() {
	//data, err := questions.ZsQuestionHttp([]string{"178658"})
	data, err := questions.EbagQuestionHttp([]string{"1910220052321159"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
