package logz

import (
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
	"log"

	"encoding/json"
	"strings"
	"time"
)

const (
	INFO  = "info"
	DEBUG = "debug"
	WARN  = "warn"
	ERROR = "error"
)

const (
	RES_OK = "请求成功"
)

type requestData struct {
	Ip         string      `json:"ip"`
	Scheme     string      `json:"scheme"`
	Url        string      `json:"url"`
	Method     string      `json:"method"`
	Form       interface{} `json:"form"`
	Body       interface{} `json:"body"`
	Header     interface{} `json:"header"`
	RemoteAddr string      `json:"remoteAddr"`
}

type logData struct {
	Time    time.Time   `json:"time"`
	Level   string      `json:"level,default=info,options=[debug,info,error,warn]"`
	Data    interface{} `json:"content,omitempty"`
	Request requestData `json:"request,omitempty"`
}

type logConfig struct {
	Level string
}

var logCfg logConfig

func LogConfig(cfg beegoConfig.Configer) {
	logCfg.Level = "debug"
}

func getLogLevel(level string) bool {
	levelList := strings.Split(logCfg.Level, ",")
	for _, v := range levelList {
		if v == level {
			return true
		}
	}
	return false
}

func printLog(data interface{}, level string) {
	if !getLogLevel(level) {
		return
	}
	logData := logData{
		Time:  time.Now(),
		Level: level,
		Data:  data,
	}
	marshal, _ := json.Marshal(logData)
	str := string(marshal)
	log.Println(str)
}

func Info(data interface{}) {
	printLog(data, INFO)
}

func Warn(data interface{}) {
	printLog(data, WARN)
}

func Debug(data interface{}) {
	printLog(data, DEBUG)
}

func Error(data interface{}) {
	printLog(data, ERROR)
}

func getRequest(beegoCtx *context.Context) requestData {
	req := beegoCtx.Request
	return requestData{
		Ip:         beegoCtx.Input.IP(),
		Scheme:     req.Proto,
		Url:        req.RequestURI,
		Method:     req.Method,
		Form:       req.Form,
		Body:       req.Body,
		Header:     req.Header,
		RemoteAddr: req.RemoteAddr,
	}
}

func printRequest(data interface{}, request requestData, level string) {
	if !getLogLevel(level) {
		return
	}
	logData := logData{
		Time:    time.Now(),
		Level:   level,
		Data:    data,
		Request: request,
	}
	marshal, _ := json.Marshal(logData)
	str := string(marshal)
	log.Println(str)
}

func RequestError(dataErr interface{}, beegoCtx *context.Context) {
	request := getRequest(beegoCtx)
	printRequest(dataErr, request, ERROR)
}

func RequestSucceed(beegoCtx *context.Context) {
	request := getRequest(beegoCtx)
	printRequest(RES_OK, request, INFO)
}

func Request(data interface{}, beegoCtx *context.Context) {
	request := getRequest(beegoCtx)
	if data != nil {
		printRequest(data, request, ERROR)
	} else {
		printRequest(RES_OK, request, INFO)
	}
}
