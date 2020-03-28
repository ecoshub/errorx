package errorx

import (
	"fmt"
)

// Error is main struct of errorx package
type Error struct {
	Header string
	Info   string
	Code   int
	Inner  []error
}

func New(Header, Info string, Code int) *Error {
	return &Error{Header: Header, Info: Info, Code: Code}
}

func (e *Error) Link(err error) *Error {
	e.Inner = append(e.Inner, err)
	return e
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	lene := len(e.Inner)
	if lene == 0 {
		return fmt.Sprintf("%v: %v, Error Code:%03d ", e.Header, e.Info, e.Code)
	}
	str := fmt.Sprintf("%v: %v, Error Code:%03d", e.Header, e.Info, e.Code)
	str += " Linked Errors: "
	for i := 0; i < lene-1; i++ {
		err := e.Inner[i]
		str += fmt.Sprintf("\t- %v\t", err.Error())
	}
	err := e.Inner[lene-1]
	str += fmt.Sprintf("\t- %v", err.Error())
	return str
}
