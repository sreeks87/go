package domain

import "fmt"

type ValidationError struct {
	ErrCode   int
	ErrorDesc string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("error : %d - %s", v.ErrCode, v.ErrorDesc)
}

func (v ValidationError) Code() int {
	return v.ErrCode
}
