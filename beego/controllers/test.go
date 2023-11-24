package controllers

import (
	"StandardProject/beego/global"
	"StandardProject/beego/services"
	"StandardProject/beego/types/request"
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
	//test1, err := testService.Test3(params.Id)
	//if err != nil {
	//	c.OutputError(err)
	//	return
	//}
	//test1, err := testService.Test2(params.Id, params.Name)
	//if err != nil {
	//	c.OutputError(err)
	//	return
	//}
	test1, err := testService.Test1()
	if err != nil {
		c.OutputError(err)
		return
	}
	//c.OutputSuccess(test1)
	c.Output(global.OK, "vsvs", test1, nil)
}
