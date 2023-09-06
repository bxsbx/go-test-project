package controllers

import (
	"TestGeneral/services"
	"TestGeneral/types/request"
	"errors"
	"github.com/astaxie/beego/validation"
)

type TestGroupController struct {
	BaseController
}

// @Summary	测试分组2
// @Tags		TestGroup
// @Produce	json
// @Param		id		formData		int		false	"id主键"
// @Param		name	formData	string	true	"名称"
// @Param		is_ok	formData		bool	false	" "
// @Param		money	formData		float64	true	" "
// @Response	200		{object}	response.Response{data=response.TestFuncResp}
// @Router		/v1/web/teacher/test/get [get]
func (c *TestGroupController) TestFunc() {
	var req request.TestFuncReq
	if err := c.ParseForm(&req); nil != err {
		c.OutputError(err)
		return
	}
	v := validation.Validation{}
	if ok, err := v.Valid(&req); !ok || nil != err {
		c.OutputError(errors.New("校验不通过"))
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
