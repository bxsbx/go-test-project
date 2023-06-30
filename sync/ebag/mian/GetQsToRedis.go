package main

import (
	"StandardProject/sync/ebag/questions"
	"StandardProject/sync/ebag/redis"
	"log"
)

func main() {

	qsIds, err := questions.GetQsIdFromBigData(MOMENT)
	if err != nil {
		log.Fatal(err)
	}
	zsQsIds, ebagQsIds := questions.GetZsAndEbagQsIds(qsIds)
	zsQsIds = questions.FilerExistIds(zsQsIds, redis.ZS_QUESTIONS)
	ebagQsIds = questions.FilerExistIds(ebagQsIds, redis.EBAG_QUESTIONS)

	go questions.GetZsQuestions(zsQsIds, COROUTINE_NUM)

	go questions.GetEbagQuestions(ebagQsIds, COROUTINE_NUM)
}
