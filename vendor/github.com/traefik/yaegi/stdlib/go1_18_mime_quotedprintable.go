// Code generated by 'yaegi extract mime/quotedprintable'. DO NOT EDIT.

//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package stdlib

import (
	"mime/quotedprintable"
	"reflect"
)

func init() {
	Symbols["mime/quotedprintable/quotedprintable"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewReader": reflect.ValueOf(quotedprintable.NewReader),
		"NewWriter": reflect.ValueOf(quotedprintable.NewWriter),

		// type definitions
		"Reader": reflect.ValueOf((*quotedprintable.Reader)(nil)),
		"Writer": reflect.ValueOf((*quotedprintable.Writer)(nil)),
	}
}
