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

package vm

import (
	"github.com/bytecodealliance/wasmtime-go/v3"

	"github.com/onflow/cadence/runtime/common"
)

var ExternFunctionsWrappers = func(runtime *WASMRuntime) []wasmtime.AsExtern {
	return []wasmtime.AsExtern{
		wasmtime.WrapFunc(runtime.store, NewStruct),
		wasmtime.WrapFunc(runtime.store, StringLoadExternFunc),
	}
}

var ExternFunctions = func(runtime *WASMRuntime) map[string]interface{} {
	return map[string]interface{}{
		"new_struct":           NewStruct,
		"new_composite_value":  NewCompositeValueExternFunc(runtime),
		"string_const":         StringLoadExternFunc,
		"new_address_location": NewAddressLocationFromHex,
	}
}

func StringLoadExternFunc(caller *wasmtime.Caller, index, len int32) string {
	// Should always read from caller's memory/data
	// TODO: any better way to access 'data' section?
	memory := caller.GetExport("memory").Memory()
	data := memory.UnsafeData(caller)
	return string(data[index : index+len])
}

func NewCompositeValueExternFunc(runtime *WASMRuntime) func(
	location common.Location,
	qualifiedIdentifier string,
	kind int32,
) *CompositeValue {
	return func(location common.Location, qualifiedIdentifier string, kind int32) *CompositeValue {

		// TODO: validate
		compositeKind := common.CompositeKind(kind)

		// always created on stack
		address := common.Address{}

		return NewCompositeValue(
			runtime,
			location,
			qualifiedIdentifier,
			compositeKind,
			address,
		)
	}
}
