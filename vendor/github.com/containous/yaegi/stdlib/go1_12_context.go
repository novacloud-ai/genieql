// Code generated by 'goexports context'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"context"
	"reflect"
	"time"
)

func init() {
	Symbols["context"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Background":       reflect.ValueOf(context.Background),
		"Canceled":         reflect.ValueOf(&context.Canceled).Elem(),
		"DeadlineExceeded": reflect.ValueOf(&context.DeadlineExceeded).Elem(),
		"TODO":             reflect.ValueOf(context.TODO),
		"WithCancel":       reflect.ValueOf(context.WithCancel),
		"WithDeadline":     reflect.ValueOf(context.WithDeadline),
		"WithTimeout":      reflect.ValueOf(context.WithTimeout),
		"WithValue":        reflect.ValueOf(context.WithValue),

		// type definitions
		"CancelFunc": reflect.ValueOf((*context.CancelFunc)(nil)),
		"Context":    reflect.ValueOf((*context.Context)(nil)),

		// interface wrapper definitions
		"_Context": reflect.ValueOf((*_context_Context)(nil)),
	}
}

// _context_Context is an interface wrapper for Context type
type _context_Context struct {
	WDeadline func() (deadline time.Time, ok bool)
	WDone     func() <-chan struct{}
	WErr      func() error
	WValue    func(key interface{}) interface{}
}

func (W _context_Context) Deadline() (deadline time.Time, ok bool) { return W.WDeadline() }
func (W _context_Context) Done() <-chan struct{}                   { return W.WDone() }
func (W _context_Context) Err() error                              { return W.WErr() }
func (W _context_Context) Value(key interface{}) interface{}       { return W.WValue(key) }
