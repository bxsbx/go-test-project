package group

import (
	"TestGeneral/controllers"
	"github.com/astaxie/beego"
)

func TestGroup() {

	//测试分组2
	beego.Router("/v1/web/teacher/test/get", &controllers.TestGroupController{}, "get:TestFunc")
	// router general tag
}
