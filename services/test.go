package services

import (
	"StandardProject/models"
	"StandardProject/types/db"
	"context"
	"time"
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

func (c *testService) Test3() (string, error) {
	time.Sleep(time.Second * 2)
	return "sleep", nil
}
