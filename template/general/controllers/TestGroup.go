package controllers

import (
	"TestGeneral/services"
	"TestGeneral/types/request"
)

type TestGroupController struct {
	BaseController
}

func (c *TestGroupController) TestFunc() {
	var req request.TestFuncReq
	err := c.FormValidate(&req)
	if err != nil {
		c.OutputError(err)
		return
	}
	testFuncService := services.NewTestGroupService(c.AppCtx)
	data, err := testFuncService.TestFunc()
	if err != nil {
		c.OutputError(err)
		return
	}
	c.OutputSuccess(data)
}

// @Summary	测试分组2
// @Tags		TestGroup
// @Produce	json
// @Param		id1		formData	int		false	"id主键1"
// @Param		name1	formData	string	true	"名称1"
// @Param		is_ok1	formData	bool	false	"是否ok"
// @Param		money1	formData	float64	true	"钱"
// @Response	200		{object}	response.Response{data=response.TestFunc2Resp}
// @Router		/v1/web/teacher/test/func [post]
func (c *TestGroupController) TestFunc2() {
	var req request.TestFunc2Req
	err := c.FormValidate(&req)
	if err != nil {
		c.OutputError(err)
		return
	}

	testFunc2Service := services.NewTestGroupService(c.AppCtx)
	data, err := testFunc2Service.TestFunc2()
	if err != nil {
		c.OutputError(err)
		return
	}
	c.OutputSuccess(data)
}
