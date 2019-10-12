// Code generated by 'goexports net/http/httptrace'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"net/http/httptrace"
	"reflect"
)

func init() {
	Symbols["net/http/httptrace"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"ContextClientTrace": reflect.ValueOf(httptrace.ContextClientTrace),
		"WithClientTrace":    reflect.ValueOf(httptrace.WithClientTrace),

		// type definitions
		"ClientTrace":      reflect.ValueOf((*httptrace.ClientTrace)(nil)),
		"DNSDoneInfo":      reflect.ValueOf((*httptrace.DNSDoneInfo)(nil)),
		"DNSStartInfo":     reflect.ValueOf((*httptrace.DNSStartInfo)(nil)),
		"GotConnInfo":      reflect.ValueOf((*httptrace.GotConnInfo)(nil)),
		"WroteRequestInfo": reflect.ValueOf((*httptrace.WroteRequestInfo)(nil)),
	}
}