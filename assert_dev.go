//go:build !release

package assure

import "fmt"

type FailHandler func(msg string)

// / Default fail handler that will panic with a message.
func PanicHandler() FailHandler {
	return func(msg string) {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}

// / Sets new fail handler as default to be used by assertions.
// / Important: Not safe in a concurrent environment.
func SetDefaultFailHandler(newHandler FailHandler) {
	defaultContext = NewContext(newHandler)
}

var defaultContext = NewContext(PanicHandler())

func Assert(condition func() bool, msg string) {
	defaultContext.Assert(condition, msg)
}

func Equal(a any, b any, msg string) {
	defaultContext.Equal(a, b, msg)
}

func NotEqual(a any, b any, msg string) {
	defaultContext.NotEqual(a, b, msg)
}

func Nil(a any, msg string) {
	defaultContext.Nil(a, msg)
}

func NotNil(a any, msg string) {
	defaultContext.NotNil(a, msg)
}

func Err(err error, msg string) {
	defaultContext.Err(err, msg)
}

func NoErr(err error, msg string) {
	defaultContext.NoErr(err, msg)
}

type Context struct {
	failHandler FailHandler
}

func NewContext(failHandler FailHandler) Context {
	return Context{failHandler}
}

func (ctx *Context) Assert(condition func() bool, msg string) {
	if !condition() {
		ctx.failHandler(msg)
	}
}

func (ctx *Context) Equal(a any, b any, msg string) {
	if a != b {
		ctx.failHandler(msg)
	}
}

func (ctx *Context) NotEqual(a any, b any, msg string) {
	if a == b {
		ctx.failHandler(msg)
	}
}

func (ctx *Context) Nil(a any, msg string) {
	if a != nil {
		ctx.failHandler(msg)
	}
}

func (ctx *Context) NotNil(a any, msg string) {
	if a == nil {
		ctx.failHandler(msg)
	}
}

func (ctx *Context) Err(err error, msg string) {
	if err == nil {
		ctx.failHandler(msg)
	}
}

func (ctx *Context) NoErr(err error, msg string) {
	if err != nil {
		ctx.failHandler(msg)
	}
}
