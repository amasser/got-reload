// Code generated by 'yaegi extract gioui.org/gpu/gl'. DO NOT EDIT.

package main

import (
	"gioui.org/gpu/gl"
	"reflect"
)

func init() {
	Symbols["gioui.org/gpu/gl"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewBackend": reflect.ValueOf(gl.NewBackend),

		// type definitions
		"Backend": reflect.ValueOf((*gl.Backend)(nil)),
		"Context": reflect.ValueOf((*gl.Context)(nil)),

		// interface wrapper definitions
		"_Context": reflect.ValueOf((*_gioui_org_gpu_gl_Context)(nil)),
	}
}

// _gioui_org_gpu_gl_Context is an interface wrapper for Context type
type _gioui_org_gpu_gl_Context struct {
}
