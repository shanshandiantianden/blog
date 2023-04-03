package utils

const (
	SUCCESS              = 200
	ERROR                = 500
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USERNAME_NOT   = 1003
	ERROR_CATEGORY       = 1007
	ERROR_ART_NULL       = 1008
	ERROR_CAART_NULL     = 1009
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "ERROR",
	ERROR_USERNAME_USED:  "用户存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USERNAME_NOT:   "用户不存在",
	ERROR_CATEGORY:       "分类存在",
	ERROR_ART_NULL:       "文章不存在",
	ERROR_CAART_NULL:     "分类文章不存在",
}

func GetErrMsg(code int) string {

	return codeMsg[code]
}
