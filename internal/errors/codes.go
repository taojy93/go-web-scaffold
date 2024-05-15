package errors

// 定义错误码
const (
	CodeSuccess       = 0
	CodeBadRequest    = 400
	CodeInternalError = 500
	CodeNotFound      = 404
	CodeInvalidParams = 10001
	CodeDatabaseError = 10002
	CodeUnauthorized  = 401
)

// 定义错误消息
var codeMsg = map[int]string{
	CodeSuccess:       "Success",
	CodeBadRequest:    "Bad Request",
	CodeInternalError: "Internal Server Error",
	CodeNotFound:      "Not Found",
	CodeInvalidParams: "Invalid Parameters",
	CodeDatabaseError: "Database Error",
	CodeUnauthorized:  "Unauthorized",
}

// 获取错误消息
func GetMsg(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	}
	return codeMsg[CodeInternalError]
}
