package main

import (
	"StandardProject/common/util"
	"fmt"
	"log"
	"strings"
	"text/template"
)

func main() {
	projectName := "TestGeneral"
	tableName := "student_mongo"

	templatePath := "H:\\GoProject\\StandardProject\\template\\beego_file\\mongo.tmpl"
	filePath := "H:\\GoProject\\StandardProject\\template\\general\\models\\" + tableName + ".go"

	mongoTemplate := MongoTemplate{
		ProjectName:    projectName,
		UpperTableName: util.HumpNaming(tableName),
	}
	mongoTemplate.LowerTableName = util.FirstLower(mongoTemplate.UpperTableName)

	exist, err := util.FileIsExist(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		t := template.Must(template.ParseFiles(templatePath))
		var builder strings.Builder
		err = t.Execute(&builder, mongoTemplate)
		if err != nil {
			log.Fatal(err)
		}
		err = util.WriteToFile(filePath, builder.String())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("创建成功")
	}
}

type MongoTemplate struct {
	ProjectName    string
	UpperTableName string
	LowerTableName string
}
