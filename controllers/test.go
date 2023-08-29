package controllers

import (
	"StandardProject/global"
	"StandardProject/services"
	"StandardProject/types/request"
)

type TestController struct {
	BaseController
}

func (c *TestController) Test1() {
	var params request.Test1Params
	err := c.FormValidate(&params)
	if err != nil {
		c.OutputError(err)
		return
	}
	testService := services.NewTestService(c.AppCtx)
	test1, err := testService.Test3()
	if err != nil {
		c.OutputError(err)
		return
	}
	//c.OutputSuccess(test1)
	c.Output(global.OK, "vsvs", test1, nil)
}
