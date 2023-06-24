package auth

import (
	"StandardProject/client"
	"github.com/astaxie/beego/context"
)

const (
	AUTH     = "auth"
	AUTH_ERR = "authErr"

	TEACHER = "teacher"
	STUDENT = "student"
	SERVER  = "server"
)

func NoAuthToken(ctx *context.Context) {
	ctx.Input.SetData(AUTH, true)
}

func authToken(ctx *context.Context, userType, userId string) {
	pass := ctx.Input.GetData(AUTH)
	if pass == nil || !pass.(bool) {
		curUserType := ctx.Request.Header.Get("userType")
		if curUserType == userType {
			err := client.AuthToken("token", "1", userType, userId)
			if err != nil {
				ctx.Input.SetData(AUTH_ERR, err)
				ctx.Input.SetData(AUTH, false)
			} else {
				ctx.Input.SetData(AUTH, true)
			}
		} else {
			ctx.Input.SetData(AUTH, false)
		}
	}
}

func AuthTeacherToken(ctx *context.Context) {
	authToken(ctx, TEACHER, "3434")
}

func AuthStudentToken(ctx *context.Context) {
	authToken(ctx, STUDENT, "2572")
}

func AuthServerToken(ctx *context.Context) {
	authToken(ctx, SERVER, "")
}
