package main

import (
	"StandardProject/sync/ebag/mysql"
	"StandardProject/sync/ebag/questions"
	"fmt"
)

func main() {
	//creatTime, _ := time.Parse(time.DateTime, "2022-05-11 08:00:00")
	//question := questions.StudentWrongQuestions{
	//	StudentID:   "0-2019-1-eb9b3075-54d0-493b-ae8d-0f6b39d52400",
	//	QuestionID:  "253807668",
	//	CorrectNum:  0,
	//	ErrorNum:    2,
	//	ScoreRate:   0,
	//	MasterLevel: 0,
	//	SchoolID:    192158,
	//	ClassID:     "0-2019-1-4e45f559-ec40-4cf5-8583-c4d1a4830ea0",
	//	SubjectID:   2,
	//	From:        4,
	//	WrongTag:    0,
	//	EbagTag:     3,
	//	IsCollect:   -1,
	//	IsWrong:     -1,
	//	CreatedAt:   creatTime,
	//	UpdatedAt:   time.Now(),
	//}
	//err := questions.CreateWrongQs(question, mysql.MysqlDB)
	err := questions.UpdateWrongQs("0-2019-1-eb9b3075-54d0-493b-ae8d-0f6b39d52400", "253807668", mysql.MysqlDB)
	if err != nil {
		fmt.Println(err)
	}
}
