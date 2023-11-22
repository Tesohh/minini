package action

import "fmt"

var (
	ErrFieldEmpty    = fmt.Errorf("required field is empty")
	ErrUserNotFound  = fmt.Errorf("user not found")
	ErrWrongPassword = fmt.Errorf("wrong password")
)
