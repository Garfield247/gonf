package e

const (
	StatusOK        = 0
	InvaldToken     = 101
	FaildParseToken = 102
	TokenTimeout    = 103
	StatusNotFound  = 404
	UnknowError     = 999
)

var statusText = map[int]string{
	StatusOK:        "OK",
	StatusNotFound:  "目标不存在或已删除",
	InvaldToken:     "无效TOKEN",
	FaildParseToken: "解析TOKEN失败",
	TokenTimeout:    "TOKEN超时",
	UnknowError:     "未预期到的错误",
}

func StatusText(code int) string {
	return statusText[code]
}

func GetStatusText() map[int]string {
	return statusText
}
