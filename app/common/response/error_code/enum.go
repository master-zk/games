package error_code

const (
	Success = 0 // 未登录

	NotLogin = 1000 // 未登录

	AccountNotActive = 1010 // 账户未激活
	AccountDisabled  = 1011 // 账户禁用

	TokenExpire = 2000 // token过期

	MessageError = 4000 // 自定义错误信息
)
