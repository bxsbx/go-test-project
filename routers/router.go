package routers

import (
	"StandardProject/common/errorz"
	"StandardProject/routers/auth"
	"StandardProject/routers/children"
	"StandardProject/types/response"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)

/*
路由规则，版本号/调用角色/平台/细化1（表名）/具体作用/...
*/

func init() {
	//后端允许跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	// 判断登录是否成功
	beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context) {
		pass := ctx.Input.GetData(auth.AUTH)
		if pass != nil && !pass.(bool) {
			errData := ctx.Input.GetData(auth.AUTH_ERR)
			if errData != nil {
				err := errData.(error)
				code, msg := errorz.GlobalError(err)
				resp := response.Response{Code: code, Msg: msg}
				ctx.Output.JSON(resp, false, false)
			}
		}
	})

	children.Teacher()
	children.Student()
	children.Server()
	children.Other()
}
