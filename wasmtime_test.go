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
	"github.com/onflow/cadence/runtime/common"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bytecodealliance/wasmtime-go/v3"
)

func TestRecursiveFib_wasmtime(t *testing.T) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(t, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(t, err)

	runtime := vm.NewWASMRuntime(store)

	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(t, err)

	fib := instance.GetFunc(store, "recursive_fib")

	result, err := fib.Call(store, 14)
	require.NoError(t, err)

	assert.Equal(t, int32(377), result)
}

func TestImperativeFib_wasmtime(t *testing.T) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(t, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(t, err)

	runtime := vm.NewWASMRuntime(store)
	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(t, err)

	fib := instance.GetFunc(store, "imperative_fib")

	result, err := fib.Call(store, 14)
	require.NoError(t, err)

	assert.Equal(t, int32(377), result)
}

func BenchmarkRecursiveFib_wasmtime(b *testing.B) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(b, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(b, err)

	runtime := vm.NewWASMRuntime(store)
	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fib := instance.GetFunc(store, "recursive_fib")
		_, _ = fib.Call(store, 14)
	}
}

func BenchmarkImperativeFib_wasmtime(b *testing.B) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(b, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(b, err)

	runtime := vm.NewWASMRuntime(store)
	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fib := instance.GetFunc(store, "imperative_fib")
		_, _ = fib.Call(store, 14)
	}
}

func BenchmarkModuleLoading_wasmtime(b *testing.B) {

	wasmBytes, _ := os.ReadFile("module.wasm")
	store := wasmtime.NewStore(wasmtime.NewEngine())

	b.ReportAllocs()
	b.ResetTimer()

	runtime := vm.NewWASMRuntime(store)

	for i := 0; i < b.N; i++ {
		module, _ := wasmtime.NewModule(store.Engine, wasmBytes)
		_, _ = wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	}
}

func TestExternFunction_wasmtime(t *testing.T) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(t, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(t, err)

	runtime := vm.NewWASMRuntime(store)
	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(t, err)

	function := instance.GetFunc(store, "create_struct_simple")

	result, err := function.Call(store)
	require.NoError(t, err)

	assert.Equal(t, vm.Struct{Name: "Foo"}, result)
}

func BenchmarkExternFunction_wasmtime(b *testing.B) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(b, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(b, err)

	runtime := vm.NewWASMRuntime(store)
	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	function := instance.GetFunc(store, "create_struct")
	for i := 0; i < b.N; i++ {
		_, _ = function.Call(store)
	}
}

func BenchmarkEmptyFunction_wasmtime(b *testing.B) {

	wasmBytes, err := os.ReadFile("module.wasm")
	require.NoError(b, err)

	store := wasmtime.NewStore(wasmtime.NewEngine())

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	require.NoError(b, err)

	runtime := vm.NewWASMRuntime(store)
	instance, err := wasmtime.NewInstance(store, module, vm.ExternFunctionsWrappers(runtime))
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	function := instance.GetFunc(store, "empty_function")
	for i := 0; i < b.N; i++ {
		_, _ = function.Call(store)
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
		vm.NewStruct()
	}
}

func TestNewCompositeValue(t *testing.T) {

	store := wasmtime.NewStore(wasmtime.NewEngine())
	runtime := vm.NewWASMRuntime(store)
	externFuncs := vm.ExternFunctions(runtime)

	mainContract, err := wasmtime.NewModuleFromFile(store.Engine, "module1.wasm")
	require.NoError(t, err)

	importedContract, err := wasmtime.NewModuleFromFile(store.Engine, "module2.wasm")
	require.NoError(t, err)

	fields := map[string]interface{}{}
	externFuncs["set_int_member"] = func(self *vm.CompositeValue, fieldName string, value int32) {
		fields[fieldName] = value
	}
	externFuncs["set_member"] = func(self *vm.CompositeValue, fieldName string, value interface{}) {
		fields[fieldName] = value
	}

	linker := wasmtime.NewLinker(store.Engine)
	for name, function := range externFuncs {
		err = linker.DefineFunc(store, "", name, function)
		require.NoError(t, err)
	}

	// The imported module is instantiated first since it has no imports
	importedContractInstance, err := linker.Instantiate(store, importedContract)
	require.NoError(t, err)

	err = linker.DefineInstance(store, "0x2.ImportedContract", importedContractInstance)
	require.NoError(t, err)

	// Then instantiate the main module
	mainContractInstance, err := linker.Instantiate(store, mainContract)
	require.NoError(t, err)

	err = linker.DefineInstance(store, "0x1.MainContract", mainContractInstance)
	require.NoError(t, err)

	// Call the function

	function := mainContractInstance.GetFunc(store, "create_composite_value")
	result, err := function.Call(store)
	require.NoError(t, err)

	require.IsType(t, &vm.CompositeValue{}, result)
	compositeValue := result.(*vm.CompositeValue)

	assert.Equal(t, "Foo", compositeValue.QualifiedIdentifier)
	assert.Equal(t, common.CompositeKindStructure, compositeValue.Kind)

	require.Len(t, fields, 2)
	assert.Equal(t, int32(4), fields["a"])
	assert.Equal(t, "Hello", fields["b"])
}

func BenchmarkNewCompositeValue(b *testing.B) {

	store := wasmtime.NewStore(wasmtime.NewEngine())
	runtime := vm.NewWASMRuntime(store)
	externFuncs := vm.ExternFunctions(runtime)

	mainContract, err := wasmtime.NewModuleFromFile(store.Engine, "module1.wasm")
	require.NoError(b, err)

	importedContract, err := wasmtime.NewModuleFromFile(store.Engine, "module2.wasm")
	require.NoError(b, err)

	//fields := map[string]interface{}{}
	externFuncs["set_int_member"] = func(self *vm.CompositeValue, fieldName string, value int32) {
		//fields[fieldName] = value
	}
	externFuncs["set_member"] = func(self *vm.CompositeValue, fieldName string, value interface{}) {
		//fields[fieldName] = value
	}

	linker := wasmtime.NewLinker(store.Engine)
	for name, function := range externFuncs {
		err = linker.DefineFunc(store, "", name, function)
		require.NoError(b, err)
	}

	// The imported module is instantiated first since it has no imports
	importedContractInstance, err := linker.Instantiate(store, importedContract)
	require.NoError(b, err)

	err = linker.DefineInstance(store, "0x2.ImportedContract", importedContractInstance)
	require.NoError(b, err)

	// Then instantiate the main module
	mainContractInstance, err := linker.Instantiate(store, mainContract)
	require.NoError(b, err)

	err = linker.DefineInstance(store, "0x1.MainContract", mainContractInstance)
	require.NoError(b, err)

	b.ReportAllocs()
	b.ResetTimer()

	function := mainContractInstance.GetFunc(store, "benchmark_composite_value")
	for i := 0; i < b.N; i++ {
		_, _ = function.Call(store, 10)
	}
}
