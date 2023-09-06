package controllers

import (
	"TestGeneral/types/response"
	"context"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	AppCtx context.Context
	beego.Controller
}

// Prepare runs after Init before http function execution.
func (b *BaseController) Prepare() {
	b.AppCtx = context.Background()
}

func (b *BaseController) Finish() {
}

func (b *BaseController) backData(data response.Response) {
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	b.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	b.Data["json"] = data
	b.ServeJSON()
}

func (b *BaseController) OutputError(err error) {
	b.backData(response.Response{Code: 200, Msg: err.Error()})
}

func (b *BaseController) OutputSuccess(data interface{}) {
	b.backData(response.Response{Code: 200, Msg: "ok", Data: data})
}

// 解析表单数据
func (b *BaseController) FormValidate(params interface{}) error {
	if err := b.ParseForm(params); nil != err {
		return err
	}
	v := validation.Validation{}
	if ok, err := v.Valid(params); !ok || nil != err {
		return errors.New("校验不通过")
	}
	return nil
}

// 解析json数据
func (b *BaseController) JsonValidate(params interface{}) error {
	body := b.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, params); nil != err {
		return err
	}
	v := validation.Validation{}
	if ok, err := v.Valid(params); !ok || nil != err {
		return errors.New("校验不通过")
	}
	return nil
}
