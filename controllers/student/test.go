package controllers

import (
	"StandardProject/controllers"
	"StandardProject/global"
	services "StandardProject/services/student"
	reqType "StandardProject/types/reqType/student"
)

type TestController struct {
	controllers.BaseController
}

func (c *TestController) Test1() {
	var params reqType.Test1Params
	err := c.FormValidate(&params)
	if err != nil {
		c.OutputError(err)
		return
	}
	testService := services.NewTestService(c.AppCtx)
	test1, err := testService.Test2()
	if err != nil {
		c.OutputError(err)
		return
	}
	//c.OutputSuccess(test1)
	c.Output(global.OK, "vsvs", test1, nil)
}
