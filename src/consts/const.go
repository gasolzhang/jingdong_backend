package consts

const (
	//成功返回数据
	ErrorCodeOk int = 600
	//有错误，但是客户端不用特殊处理，服务端需要根据错误信息确定具体问题
	ErrorCodeServer int = 601
	//有错误，并且需要客户端特殊处理
	ErrorCodeParameter     int = 602
	ErrorCodeTokenError    int = 603
)
