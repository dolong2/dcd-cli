package err

// CmdError 커스텀 에러 타입 정의
type CmdError struct {
	Code    int
	Message string
}

func (e *CmdError) Error() string {
	return e.Message
}

func NewCmdError(code int, message string) error {
	return &CmdError{
		Code:    code,
		Message: message,
	}
}
