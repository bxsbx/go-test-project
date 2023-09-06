package main

import (
	"StandardProject/common/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"
)

const (
	ProjectName     = "TestGeneral"
	ControllersPath = "H:\\GoProject\\StandardProject\\template\\general\\controllers"
	ServicesPath    = "H:\\GoProject\\StandardProject\\template\\general\\services"
	RouterGroupPath = "H:\\GoProject\\StandardProject\\template\\general\\router\\group"

	TypePath = "H:\\GoProject\\StandardProject\\template\\general\\types\\request\\test.go"

	ValidateForm = "FormValidate"
	ValidateJson = "JsonValidate"

	Produce = "json"

	ControllersTmpl = "H:\\GoProject\\StandardProject\\template\\beego_file\\controller.tmpl"
	ServicesTmpl    = "H:\\GoProject\\StandardProject\\template\\beego_file\\service.tmpl"
	RouterTmpl      = "H:\\GoProject\\StandardProject\\template\\beego_file\\router.tmpl"
)

var ParamTypeMap = map[string]string{
	"path":  "path",
	"query": "query",
	"form":  "formData",
	"json":  "body",
}

func main() {
	myTemplate := myTemplate{
		ProjectName:  ProjectName,
		GroupName:    "TestGroup",
		FunName:      "TestFunc",
		ReqType:      "TestFuncReq",
		Router:       "/v1/web/teacher/test/get",
		Method:       strings.ToLower(http.MethodGet),
		ValidateType: ValidateForm,
	}

	swagger := Swagger{
		Summary: "测试分组2",
		Tags:    myTemplate.GroupName,
		Produce: Produce,
	}
	swagger.Params = GetParamsFromReqType(myTemplate.ReqType)
	if swagger.Params == nil {
		log.Fatal("请求类型不存在")
	}
	myTemplate.Swagger = swagger

	ControllersFilePath := ControllersPath + "\\" + myTemplate.GroupName + ".go"
	isControllerFile, _ := util.FileIsExist(ControllersFilePath)
	myTemplate.IsControllerFile = !isControllerFile

	ServicesFilePath := ServicesPath + "\\" + myTemplate.GroupName + ".go"
	isServiceFile, _ := util.FileIsExist(ServicesFilePath)
	myTemplate.IsServiceFile = !isServiceFile

	RouterFilePath := RouterGroupPath + "\\" + myTemplate.GroupName + ".go"
	isRouterFile, _ := util.FileIsExist(RouterFilePath)
	myTemplate.IsRouterFile = !isRouterFile

	myTemplate.LowerFunName = util.FirstLower(myTemplate.FunName)

	funcName := func(line string) bool {
		if len(line) > 4 && line[:4] == "func" && strings.Contains(line, myTemplate.FunName) {
			return true
		}
		return false
	}

	routerPath := func(line string) bool {
		if strings.Contains(line, myTemplate.Router) {
			return true
		}
		return false
	}
	GeneralContentToFile(ControllersTmpl, ControllersFilePath, "", myTemplate, funcName)
	GeneralContentToFile(ServicesTmpl, ServicesFilePath, "", myTemplate, funcName)
	GeneralContentToFile(RouterTmpl, RouterFilePath, "router general tag", myTemplate, routerPath)

	//swag注释格式化
	exec.Command("swag", "fmt", "-d", ControllersFilePath).Run()
	err := exec.Command("swag", "init", "-dir", "H:\\GoProject\\StandardProject\\template\\general").Run()
	if err != nil {
		fmt.Println(err)
	}
}

type Swagger struct {
	Summary string
	Tags    string
	Produce string
	Params  []Params
}
type Params struct {
	Name        string
	ParamType   string
	DataType    string
	IsNeed      bool
	Description string
}

type myTemplate struct {
	ProjectName      string
	GroupName        string
	FunName          string
	LowerFunName     string
	ReqType          string
	Swagger          Swagger
	Router           string
	Method           string
	ValidateType     string
	IsRouterFile     bool
	IsControllerFile bool
	IsServiceFile    bool
}

// 从请求类型中获取Swagger信息
func GetParamsFromReqType(reqType string) []Params {
	reqTypeStr := "type " + reqType + " struct {"
	bytes, err := os.ReadFile(TypePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bytes), "\n")

	var list []Params
	flag := false
	for _, line := range lines {
		if strings.Contains(line, reqTypeStr) {
			flag = true
		} else if len(line) > 0 && line[:1] == "}" && flag {
			break
		} else if flag {
			var params Params
			strs := regexp.MustCompile("\\s+").Split(strings.Trim(line, " |\t"), -1)
			params.DataType = strs[1]
			if strings.Contains(line, "Required") {
				params.IsNeed = true
			}

			tags := strings.Trim(regexp.MustCompile("`.*`").FindString(line), "`")
			params.Name = strings.Trim(regexp.MustCompile("\"\\w*\"").FindString(tags), "\"")
			params.ParamType = ParamTypeMap[strings.Split(tags, ":")[0]]
			if params.ParamType == "" {
				params.ParamType = ParamTypeMap["form"] // 默认form
			}
			params.Description = strings.Trim(strings.Trim(regexp.MustCompile("//.*").FindString(line), "//"), " ")
			list = append(list, params)
		}
	}
	return list
}

func GeneralContentToFile(tmpl, filePath, location string, myTemplate myTemplate, f func(l string) bool) {
	exist, err := util.FileIsExist(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if exist {
		content, err := util.ReadFileContent(filePath)
		if err != nil {
			log.Fatal(err)
		}
		//方法已存在则跳过
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if f(line) {
				return
			}
		}
	}

	t := template.Must(template.ParseFiles(tmpl))
	var builder strings.Builder
	err = t.Execute(&builder, myTemplate)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(builder.String())
	if location != "" {
		err = util.AddAtCustomLocation(filePath, builder.String(), location)
	} else {
		err = util.WriteToFile(filePath, builder.String())
	}
	if err != nil {
		log.Fatal(err)
	}
}
