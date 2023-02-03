package evtrsk

type SendError struct {
	Dong   error
	Email  error
	Wechat error
	SMS    error
	Phone  error
}

func (se SendError) Error() string {
	return ""
}
