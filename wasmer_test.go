/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"cadence_wasm_vm/vm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func TestRecursiveFib_wasmer(t *testing.T) {
	wasmBytes, err := os.ReadFile("module1.wasm")
	require.NoError(t, err)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	require.NoError(t, err)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	require.NoError(t, err)

	// Gets the `recursive_fib` exported function from the WebAssembly instance.
	fib, err := instance.Exports.GetFunction("recursive_fib")
	require.NoError(t, err)

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, err := fib(14)
	require.NoError(t, err)

	assert.Equal(t, int32(377), result)
}

func TestImperativeFib_wasmer(t *testing.T) {
	wasmBytes, err := os.ReadFile("module1.wasm")
	require.NoError(t, err)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, wasmBytes)
	require.NoError(t, err)

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	require.NoError(t, err)

	fib, err := instance.Exports.GetFunction("imperative_fib")
	require.NoError(t, err)

	result, err := fib(14)
	require.NoError(t, err)

	assert.Equal(t, int32(377), result)
}

func BenchmarkRecursiveFib_wasmer(b *testing.B) {

	wasmBytes, _ := os.ReadFile("module1.wasm")
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	b.ReportAllocs()
	b.ResetTimer()

	module, _ := wasmer.NewModule(store, wasmBytes)
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	for i := 0; i < b.N; i++ {
		fib, _ := instance.Exports.GetFunction("recursive_fib")
		_, _ = fib(14)
	}
}

func BenchmarkImperativeFib_wasmer(b *testing.B) {

	wasmBytes, _ := os.ReadFile("module1.wasm")
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fib, _ := instance.Exports.GetFunction("imperative_fib")
		_, _ = fib(14)
	}
}

func BenchmarkModuleLoading_wasmer(b *testing.B) {

	wasmBytes, _ := os.ReadFile("module1.wasm")
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		module, _ := wasmer.NewModule(store, wasmBytes)
		importObject := wasmer.NewImportObject()
		_, _ = wasmer.NewInstance(module, importObject)
	}
}

func TestExternFunction_wasmer(t *testing.T) {

	wasmBytes, _ := os.ReadFile("module1.wasm")
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)
	importObject := wasmer.NewImportObject()
	importObject.Register("", map[string]wasmer.IntoExtern{
		"new_struct": wasmer.NewFunction(
			store,
			wasmer.NewFunctionType(
				[]*wasmer.ValueType{},
				wasmer.NewValueTypes(wasmer.AnyRef),
			),
			func(values []wasmer.Value) ([]wasmer.Value, error) {
				return []wasmer.Value{
					// Unfortunately wasmer doesn't support non-primitive values yet.
					// https://github.com/wasmerio/wasmer-go/issues/330
					wasmer.NewValue(vm.NewStruct(), wasmer.AnyRef),
				}, nil
			}),
	})

	instance, _ := wasmer.NewInstance(module, importObject)

	function, _ := instance.Exports.GetFunction("create_struct_simple")
	result, err := function()
	require.NoError(t, err)
	assert.Equal(t, vm.Struct{Name: "Foo"}, result)
}
