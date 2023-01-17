package main

import (
	"github.com/bytecodealliance/wasmtime-go/v3"
)

var ExternFunctions = func(store *wasmtime.Store) []wasmtime.AsExtern {
	return []wasmtime.AsExtern{
		wasmtime.WrapFunc(store, NewStruct),
	}
}
