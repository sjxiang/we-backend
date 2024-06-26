package we

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode uint32
	ErrMsg  string  
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code uint32, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}


// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}

	if errors.As(err, &Err) {
		return Err
	}

	s := ErrInternal
	s.ErrMsg = err.Error()
	return s
}
