package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000_0000, "服务内部错误")
	InvalidParams             = NewError(1000_0001, "入参错误")
	NotFound                  = NewError(1000_0002, "找不到")
	UnauthorizedAuthNotExist  = NewError(1000_0003, "鉴权失败,找不到对应的AppKey或AppSecret")
	UnauthorizedTokenError    = NewError(1000_0004, "鉴权失败,Token错误")
	UnauthorizedTokenTimeout  = NewError(1000_0005, "鉴权失败,Token超时")
	UnauthorizedTokenGenerate = NewError(1000_0006, "鉴权失败,Token生成失败")

	TooManyRequests = NewError(1000_0007, "请求过多")
)
