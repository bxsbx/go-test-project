package main

import (
	"StandardProject/common/gormdb"
	"StandardProject/common/util"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
	"text/template"
)

var DataTypeMap = map[string]string{
	"varchar":  "string",
	"text":     "string",
	"tinyint":  "int",
	"int":      "int",
	"bigint":   "int64",
	"decimal":  "float64",
	"double":   "float64",
	"float":    "float",
	"datetime": "time.Time",
	"date":     "time.Time",
}

func main() {
	projectName := "TestGeneral"
	tableName := "student"
	pre := "t_"
	names := "s"
	DeletePreTableName := strings.TrimPrefix(tableName, pre)
	templatePath := "H:\\GoProject\\StandardProject\\template\\beego_file\\mysql.tmpl"
	filePath := "H:\\GoProject\\StandardProject\\template\\general\\models\\" + DeletePreTableName + ".go"

	mysqlTemplate := MysqlTemplate{
		ProjectName:    projectName,
		ConstTableName: strings.ToUpper(DeletePreTableName),
		TableName:      tableName,
		UpperTableName: util.HumpNaming(DeletePreTableName),
	}
	mysqlTemplate.LowerTableName = util.FirstLower(mysqlTemplate.UpperTableName)
	mysqlTemplate.ListName = mysqlTemplate.LowerTableName + names

	gormdb.InitDB(nil)
	columns, err := GetTableInfo(gormdb.MyDB(), tableName)
	if err != nil {
		log.Fatal(err)
	}
	var whereList []string
	var nameTypeList []string
	var nameList []string
	for i, column := range columns {
		column.FieldName = util.HumpNaming(column.Name)
		column.Type = DataTypeMap[column.Type]
		name := util.FirstLower(column.FieldName)
		if column.PrimaryKey == "PRI" {
			nameTypeList = append(nameTypeList, name+" "+column.Type)
			nameList = append(nameList, name)
			whereList = append(whereList, name+" = ?")
		}
		columns[i] = column
	}
	mysqlTemplate.PrimaryWhere = "\"" + strings.Join(whereList, " and ") + "\", " + strings.Join(nameList, ",")
	mysqlTemplate.PrimaryParams = strings.Join(nameTypeList, ",")
	mysqlTemplate.Columns = columns

	exist, err := util.FileIsExist(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if !exist {
		t := template.Must(template.ParseFiles(templatePath))
		var builder strings.Builder
		err = t.Execute(&builder, mysqlTemplate)
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

type MysqlTemplate struct {
	ProjectName    string
	ConstTableName string
	TableName      string
	UpperTableName string
	LowerTableName string
	ListName       string
	PrimaryParams  string
	PrimaryWhere   string
	Columns        []TableColumn
}

type TableColumn struct {
	FieldName  string `gorm:"-"`
	Name       string `gorm:"column:name"`
	Type       string `gorm:"column:type"`
	Comment    string `gorm:"column:comment"`
	IsNull     string `gorm:"column:is_null"`
	PrimaryKey string `gorm:"column:primary_key"`
}

func GetTableInfo(db *gorm.DB, tableName string) (columns []TableColumn, err error) {
	query := fmt.Sprintf(`SELECT COLUMN_NAME AS name, DATA_TYPE AS type, COLUMN_COMMENT AS comment, IS_NULLABLE as is_null, COLUMN_KEY as primary_key 
					FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s' 
					ORDER BY ORDINAL_POSITION`, tableName)
	err = db.Raw(query).Scan(&columns).Error
	return
}
