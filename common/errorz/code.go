package errorz

import "net/http"

const (
	//常见错误码
	RESP_ERR         = 10001
	RESP_PARAM_ERR   = 10002
	RESP_TOKEN_ERR   = 10003
	RESP_NO_ACCESS   = 10004
	RESP_APP_NOT_ON  = 10005
	RESP_UNKNOWN_ERR = 10006

	//登录错误码
	RESP_LOGIN_UNLOG     = 12100 //未登录
	RESP_LOGIN_INCORRECT = 12101 //用户账号或者密码不正确
	RESP_LOGIN_MODIFY    = 12102 //修改密码失败
	RESP_LOGIN_FORMAT    = 12103 //用户账号格式不正确
	RESP_LOGIN_SESSION   = 12104 //创建会话状态失败
	RESP_LOGIN_EXIST     = 12105 //您的账号已在其他地方登录
	RESP_LOGIN_PARAMS    = 12106 //认证参数异常

	//序列化
	ERR_DECODE    = 30001
	ERR_UNMARSHAL = 30002
	ERR_MARSHAL   = 30003

	//http请求
	NEW_REQUEST = 30011
	REQUEST_ERR = 30012

	//数据流
	IO_READ_ERR = 30021
)

var respMsg = map[int]string{

	//常见错误码
	RESP_ERR:         "操作失败",
	RESP_PARAM_ERR:   "参数错误",
	RESP_TOKEN_ERR:   "签名认证错误",
	RESP_NO_ACCESS:   "对不起，您没有此操作权限",
	RESP_APP_NOT_ON:  "暂时未提供服务",
	RESP_UNKNOWN_ERR: "未知错误",

	//登录错误码
	RESP_LOGIN_UNLOG:     "未登录",
	RESP_LOGIN_INCORRECT: "用户账号或者密码不正确",
	RESP_LOGIN_MODIFY:    "修改密码失败",
	RESP_LOGIN_FORMAT:    "用户账号格式不正确",
	RESP_LOGIN_SESSION:   "创建会话状态失败",
	RESP_LOGIN_EXIST:     "您的账号已在其他地方登录",
	RESP_LOGIN_PARAMS:    "认证参数异常",

	//序列化
	ERR_DECODE:    "解码失败",
	ERR_UNMARSHAL: "解析失败",
	ERR_MARSHAL:   "byte数组获取失败，marshal",

	//http请求
	NEW_REQUEST: "创建请求失败",
	REQUEST_ERR: "请求失败",

	//数据流
	IO_READ_ERR: "数据流读取失败",
}

var httpCodeMap = map[int]int{
	RESP_PARAM_ERR: http.StatusBadRequest,
}
