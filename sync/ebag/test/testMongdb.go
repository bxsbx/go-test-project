package main

import (
	"StandardProject/sync/ebag/questions"
	"fmt"
	"log"
)

func main() {
	//list, err := questions.GetStuIdsFromBigDataErrSchoolId(0)
	//list, err := questions.GetQsIdFromBigData(1)
	list, err := questions.GetQsFromBigData(1, 1, 1000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(list)
}
