package __standard_lib

import (
	"errors"
	"testing"
	"time"
)

type cusError struct {
	Code string
	Msg  string
	Time time.Time
}

func NewCusError(code, msg string) error {
	return &cusError{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}

func (c *cusError) Error() string {
	return c.Code + ":" + c.Msg + ":" + c.Time.Format("2006-01-02 15:04:05")
}

func Test_Errors(t *testing.T) {
	err := errors.New("test error")
	t.Log(err)

	err = NewCusError("001", "test cus error")
	t.Log(err)
}
