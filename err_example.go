package main

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
)

type Error struct {
	Code       int
	Msg        string
	StackTrace string
}

// Error with stack info
func New(code int, msg string) *Error {
	return &Error{code, msg, stack(3)}
}
func Newf(code int, format string, args ...interface{}) error {
	return &Error{code, fmt.Sprintf(format, args), stack(3)}
}
func (e *Error) Error() string {
	return e.Msg + e.StackTrace
}
func (e *Error) Print() string {
	return e.Msg
}
func (e *Error) JsonError() string {
	err, _ := json.Marshal(e)
	return string(err)
}
func stack(skip int) string {
	stk := make([]uintptr, 32)
	str := ""
	l := runtime.Callers(skip, stk[:])
	for i := 0; i < l; i++ {
		f := runtime.FuncForPC(stk[i])
		name := f.Name()
		file, line := f.FileLine(stk[i])
		str += fmt.Sprintf("\n    %-30s [%s:%d]", name, path.Base(file), line)
	}
	return str
}

func main() {
	gg := Error{1, "nima", stack(1)}
	fmt.Println(gg.Print())
	fmt.Println(gg.Msg)
	fmt.Println(gg.Error())
}
