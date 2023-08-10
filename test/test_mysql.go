package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		Dsn("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		//Table("student").
		Run()
	fmt.Println(err)
}
