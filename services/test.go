package services

import (
	"StandardProject/models"
	"StandardProject/types/db"
	"context"
	"strconv"
)

type testService struct {
	appCtx context.Context
}

func NewTestService(appCtx context.Context) *testService {
	return &testService{appCtx: appCtx}
}

func (c *testService) Test1() (string, error) {
	return "vsasv", nil
}

func (c *testService) Test2() ([]db.Student, error) {
	studentModel := models.NewStudentModel(c.appCtx)
	return studentModel.FindAll()
}

func (c *testService) Test3(id int) (string, error) {
	//if id%2 == 0 {
	//	time.Sleep(time.Second * 2)
	//} else if id%3 == 0 {
	//	time.Sleep(time.Second * 3)
	//} else if id%4 == 0 {
	//	time.Sleep(time.Second * 4)
	//} else if id%5 == 0 {
	//	time.Sleep(time.Second * 5)
	//} else {
	//time.Sleep(time.Second * 1)
	//}

	return strconv.Itoa(id), nil
}
