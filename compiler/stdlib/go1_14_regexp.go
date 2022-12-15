// Code generated by 'goexports regexp'. DO NOT EDIT.

//go:build go1.14
// +build go1.14

package stdlib

import (
	"reflect"
	"regexp"
)

func init() {
	Symbols["regexp"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Compile":          reflect.ValueOf(regexp.Compile),
		"CompilePOSIX":     reflect.ValueOf(regexp.CompilePOSIX),
		"Match":            reflect.ValueOf(regexp.Match),
		"MatchReader":      reflect.ValueOf(regexp.MatchReader),
		"MatchString":      reflect.ValueOf(regexp.MatchString),
		"MustCompile":      reflect.ValueOf(regexp.MustCompile),
		"MustCompilePOSIX": reflect.ValueOf(regexp.MustCompilePOSIX),
		"QuoteMeta":        reflect.ValueOf(regexp.QuoteMeta),

		// type definitions
		"Regexp": reflect.ValueOf((*regexp.Regexp)(nil)),
	}
}
