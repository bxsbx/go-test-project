package main

import (
	_ "StandardProject/beego/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run(":8888")
}
