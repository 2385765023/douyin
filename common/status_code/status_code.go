//	之所以放在全局，是因为 service 层和 controller 层都有返回状态(错误)的需求
package status_code

import "fmt"

type Status struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func (err Status) Error() string {
	return fmt.Sprintf("%d %s", err.StatusCode, err.StatusMsg)
}

//	错误码定义
const (
	SuccessCode = 0
	//	用户身份相关的错误
	UserAlreadyExistErrCode       = 10000
	UserNotExistErrCode           = 10001
	UserNameOrPwdIncorrectErrCode = 10002

	UnknownErrCode = 50000
)

//	预定义状态
var (
	Success                   = NewStatus(SuccessCode, "操作成功")
	UserAlreadyExistErr       = NewStatus(UserAlreadyExistErrCode, "用户名已存在，请更换用户名后重试")
	UserNotExistErr           = NewStatus(UserNotExistErrCode, "用户不存在，请检查用户名是否正确")
	UserNameOrPwdIncorrectErr = NewStatus(UserNameOrPwdIncorrectErrCode, "登录失败，请检查用户名和密码是否正确")
	UnknownErr                = NewStatus(UnknownErrCode, "未知错误")
)

func NewStatus(code int32, msg string) Status {
	return Status{StatusCode: code, StatusMsg: msg} //据说小数据量情况下，返回值拷贝比指针要快
}

//	返回系统运行未知错误
func NewUnknownErr(err error) Status {
	return Status{StatusCode: UnknownErrCode, StatusMsg: "本错误生产环境中不应抛出，可能造成系统内部消息泄露。此处仅为了方便调试:" + err.Error()}
}
