package controllers

import (
	"StandardProject/common/errorz"
	"StandardProject/common/logz"
	"StandardProject/common/tracer"
	"StandardProject/global"
	"StandardProject/types/response"
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	AppCtx context.Context
	beego.Controller
}

// Prepare runs after Init before request function execution.
func (b *BaseController) Prepare() {
	b.AppCtx = context.Background()
	tracer.StarTracerSpan(b.AppCtx, b.Ctx.Request)
	tracer.SetTagSpan(b.AppCtx, "uid", b.GetString("uid"))
}

func (b *BaseController) Finish() {
	tracer.FinishSpan(b.AppCtx)
}

func (b *BaseController) backData(data response.Response) {
	b.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	b.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	b.Data["json"] = data
	b.Ctx.ResponseWriter.WriteHeader(errorz.GetHttpCodeWithCode(data.Code))
	b.ServeJSON()
}

func (b *BaseController) OutputError(err error) {
	errStack := errorz.GetErrorCallerList(err)
	logz.RequestError(errStack, b.Ctx)
	code, msg := errorz.GlobalError(err)
	b.backData(response.Response{Code: code, Msg: msg, ErrStack: errStack})
}

func (b *BaseController) OutputSuccess(data interface{}) {
	logz.RequestSucceed(b.Ctx)
	b.backData(response.Response{Code: global.OK, Msg: global.OK_MSG, Data: data})
}

func (b *BaseController) Output(code int, msg string, data interface{}, err error) {
	errStack := errorz.GetErrorCallerList(err)
	logz.Request(errStack, b.Ctx)
	response := response.Response{Code: code, Msg: msg}
	if data != nil {
		response.Data = data
	}
	if errStack != nil {
		response.ErrStack = errStack
	}
	b.backData(response)
}

// 解析表单数据
func (b *BaseController) FormValidate(params interface{}) error {
	if err := b.ParseForm(params); nil != err {
		return errorz.CodeError(errorz.ERR_UNMARSHAL, err)
	}
	v := validation.Validation{}
	if ok, err := v.Valid(params); !ok || nil != err {
		return errorz.CodeError(errorz.RESP_PARAM_ERR, err)
	}
	return nil
}

// 解析json数据
func (b *BaseController) JsonValidate(params interface{}) error {
	body := b.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, params); nil != err {
		return errorz.CodeError(errorz.ERR_UNMARSHAL, err)
	}
	v := validation.Validation{}
	if ok, err := v.Valid(params); !ok || nil != err {
		return errorz.CodeError(errorz.RESP_PARAM_ERR, err)
	}
	return nil
}
