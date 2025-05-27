package errorP

import "errors"

var (
	Error_PASSWORD_ERROR       = errors.New("密码错误")
	Error_ACCOUNT_NOT_FOUND    = errors.New("账号不存在")
	Error_ACCOUNT_LOCKED       = errors.New("账号被锁定")
	Error_ALREADY_EXISTS       = errors.New("已存在")
	Error_UNKNOWN_ERROR        = errors.New("未知错误")
	Error_USER_NOT_LOGIN       = errors.New("用户未登录")
	Error_LOGIN_FAILED         = errors.New("登录失败")
	Error_PASSWORD_EDIT_FAILED = errors.New("密码修改失败")
)

const (
	SUCCESS         = 1    // ok
	ERROR           = 2    // 内部错误
	UNKNOW_IDENTITY = 403  // 未知身份
	MysqlERR        = 1001 // mysql出错
	RedisERR        = 1002 // redis出错
)

var ErrMsg = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "内部错误",
	UNKNOW_IDENTITY: "未知身份",
	MysqlERR:        "mysql错误",
	RedisERR:        "redis错误",
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e CodeErrorResponse) Error() string {
	return e.Msg
}

func GetMsg(code int) string {
	return ErrMsg[code]
}
func CodeError(code int) error {
	return CodeErrorResponse{
		Code: code,
		Msg:  GetMsg(code),
	}
}
