package main

import (
	"StandardProject/sync/ebag/questions"
	"StandardProject/sync/ebag/redis"
	"log"
)

func main() {
	stuIds, err := questions.GetStuIdsFromBigDataErrSchoolId(MOMENT)
	if err != nil {
		log.Fatal(err)
	}
	stuIds = questions.FilerExistIds(stuIds, redis.STUDENT_SCHOOL)
	questions.SaveStudentToRedis(stuIds, COROUTINE_NUM)
}
