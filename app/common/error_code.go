package common

const (
	CodeOk = 0 // 成功

	CodeNotLogin = 1000 // 未登录

	CodeAccountNotActive = 1010 // 账户未激活
	CodeAccountDisabled  = 1011 // 账户禁用

	CodeTokenExpire = 2000 // token过期

	CodeMessageError   = 4000 // 自定义错误信息
	CodeNotFound       = 4004 // 找不到对象
	CodeNoRoute        = 4005 // 路由未定义
	CodeApiParamsError = 4010 // 参数异常

	CodeServerError = 5000 // 服务器异常
)

var MessageMap = map[int]string{
	CodeOk: "OK",

	CodeNotLogin: "未登录",

	CodeAccountNotActive: "账户未激活",
	CodeAccountDisabled:  "账户禁用",

	CodeTokenExpire: "token过期",

	CodeMessageError: "error",
	CodeNotFound:     "找不到对象",
	CodeNoRoute:      "路由未定义",

	CodeApiParamsError: "参数异常",

	CodeServerError: "服务器异常",
}

func GetCodeMsg(code int) string {
	return MessageMap[code]
}
