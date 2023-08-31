package errorz

import (
	"gorm.io/gorm"
)

func GlobalError(err error) (code int, msg string) {
	myErr, ok := err.(*MyError)
	temp := myErr
	for ok {
		myErr = temp
		if myErr.err != nil {
			err = myErr.err
			temp, ok = err.(*MyError)
		} else {
			break
		}
	}
	switch err {
	case gorm.ErrRecordNotFound:
		code, msg = RESP_APP_NOT_ON, GetMsgWithCode(RESP_APP_NOT_ON)
	case myErr:
		code, msg = myErr.code, myErr.msg
	default:
		code, msg = RESP_ERR, GetMsgWithCode(RESP_ERR)
	}
	return
}
