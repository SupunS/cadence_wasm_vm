package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bytecodealliance/wasmtime-go/v3"
)

func TestRecursiveFib_wasmtime(t *testing.T) {

	wasmBytes, err := os.ReadFile("fib.wasm")
	require.NoError(t, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(t, err)

	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})

	fib := instance.GetFunc(store, "recursive_fib")

	result, err := fib.Call(store, 14)
	require.NoError(t, err)

	assert.Equal(t, int32(377), result)
}

func TestImperativeFib_wasmtime(t *testing.T) {

	wasmBytes, err := os.ReadFile("fib.wasm")
	require.NoError(t, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(t, err)

	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})

	fib := instance.GetFunc(store, "imperative_fib")

	result, err := fib.Call(store, 14)
	require.NoError(t, err)

	assert.Equal(t, int32(377), result)
}

func BenchmarkRecursiveFib_wasmtime(b *testing.B) {

	wasmBytes, _ := os.ReadFile("fib.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, _ := wasmtime.NewModule(store.Engine, wasmBytes)
	instance, _ := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fib := instance.GetFunc(store, "recursive_fib")
		_, _ = fib.Call(store, 14)
	}
}

func BenchmarkImperativeFib_wasmtime(b *testing.B) {

	wasmBytes, _ := os.ReadFile("fib.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, _ := wasmtime.NewModule(store.Engine, wasmBytes)
	instance, _ := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fib := instance.GetFunc(store, "imperative_fib")
		_, _ = fib.Call(store, 14)
	}
}

func BenchmarkModuleLoading_wasmtime(b *testing.B) {

	wasmBytes, _ := os.ReadFile("fib.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		module, _ := wasmtime.NewModule(store.Engine, wasmBytes)
		_, _ = wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	}
}

func TestExternFunction_wasmtime(t *testing.T) {

	wasmBytes, err := os.ReadFile("fib.wasm")
	require.NoError(t, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(t, err)

	instance, err := wasmtime.NewInstance(store, module, ExternFunctions(store))

	fib := instance.GetFunc(store, "create_struct_simple")

	result, err := fib.Call(store)
	require.NoError(t, err)

	assert.Equal(t, Struct{name: "Foo"}, result)
}

func BenchmarkExternFunction_wasmtime(b *testing.B) {

	wasmBytes, _ := os.ReadFile("fib.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, _ := wasmtime.NewModule(store.Engine, wasmBytes)
	instance, _ := wasmtime.NewInstance(store, module, ExternFunctions(store))

	b.ReportAllocs()
	b.ResetTimer()

	fib := instance.GetFunc(store, "create_struct")
	for i := 0; i < b.N; i++ {
		_, _ = fib.Call(store)
	}
}

func BenchmarkEmptyFunction_wasmtime(b *testing.B) {

	wasmBytes, _ := os.ReadFile("fib.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())
	module, _ := wasmtime.NewModule(store.Engine, wasmBytes)
	instance, _ := wasmtime.NewInstance(store, module, ExternFunctions(store))

	b.ReportAllocs()
	b.ResetTimer()

	fib := instance.GetFunc(store, "empty_function")
	for i := 0; i < b.N; i++ {
		_, _ = fib.Call(store)
	}
}

func BenchmarkGoFunction(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		create_struct()
	}
}

func create_struct() {
	for i := 0; i < 7; i++ {
		NewStruct()
	}
}
