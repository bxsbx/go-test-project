package services

import (
	"StandardProject/models"
	"StandardProject/types/dbType"
	"context"
)

type testService struct {
	appCtx context.Context
}

func NewTestService(appCtx context.Context) *testService {
	return &testService{appCtx: appCtx}
}

func (c *testService) Test1() (string, error) {
	return "vsasv", nil
	//return "vsasv", errorz.Code(errorz.IO_READ_ERR)
}

func (c *testService) Test2() ([]dbType.Student, error) {
	studentModel := models.NewStudentModel(c.appCtx)
	return studentModel.FindAll()
}
