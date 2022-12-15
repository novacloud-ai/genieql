// Code generated by 'goexports image/gif'. DO NOT EDIT.

//go:build go1.14
// +build go1.14

package stdlib

import (
	"go/constant"
	"go/token"
	"image/gif"
	"reflect"
)

func init() {
	Symbols["image/gif"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Decode":             reflect.ValueOf(gif.Decode),
		"DecodeAll":          reflect.ValueOf(gif.DecodeAll),
		"DecodeConfig":       reflect.ValueOf(gif.DecodeConfig),
		"DisposalBackground": reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"DisposalNone":       reflect.ValueOf(constant.MakeFromLiteral("1", token.INT, 0)),
		"DisposalPrevious":   reflect.ValueOf(constant.MakeFromLiteral("3", token.INT, 0)),
		"Encode":             reflect.ValueOf(gif.Encode),
		"EncodeAll":          reflect.ValueOf(gif.EncodeAll),

		// type definitions
		"GIF":     reflect.ValueOf((*gif.GIF)(nil)),
		"Options": reflect.ValueOf((*gif.Options)(nil)),
	}
}
