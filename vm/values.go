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
	"github.com/onflow/atree"

	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/errors"
)

// For testing purpose

type Struct struct {
	Name string
}

func NewStruct() Struct {
	return Struct{
		Name: "Foo",
	}
}

// Composite value

type CompositeValue struct {
	//Destructor          FunctionValue
	Location common.Location
	//staticType          StaticType
	//Stringer            func(gauge common.MemoryGauge, value *CompositeValue, seenReferences SeenReferences) string
	//InjectedFields      map[string]Value
	//ComputedFields      map[string]ComputedField
	//NestedVariables     map[string]*Variable
	//Functions           map[string]FunctionValue
	dictionary          *atree.OrderedMap
	typeID              common.TypeID
	QualifiedIdentifier string
	Kind                common.CompositeKind
	isDestroyed         bool
}

func NewCompositeValue(
	runtime *WASMRuntime,
	location common.Location,
	qualifiedIdentifier string,
	kind common.CompositeKind,
	//fields []interpreter.CompositeField,
	address common.Address,
) *CompositeValue {

	var v *CompositeValue

	constructor := func() *atree.OrderedMap {
		dictionary, err := atree.NewMap(
			runtime.Config.Storage,
			atree.Address(address),
			atree.NewDefaultDigesterBuilder(),
			NewCompositeTypeInfo(
				location,
				qualifiedIdentifier,
				kind,
			),
		)
		if err != nil {
			panic(errors.NewExternalError(err))
		}
		return dictionary
	}

	typeInfo := NewCompositeTypeInfo(
		location,
		qualifiedIdentifier,
		kind,
	)

	v = &CompositeValue{
		dictionary:          constructor(),
		Location:            typeInfo.location,
		QualifiedIdentifier: typeInfo.qualifiedIdentifier,
		Kind:                typeInfo.kind,
	}

	//for _, field := range fields {
	//	v.SetMember(
	//		interpreter,
	//		interpreter.EmptyLocationRange,
	//		field.Name,
	//		field.Value,
	//	)
	//}

	return v
}

func NewAddressLocationFromHex(addressHex, name string) common.AddressLocation {
	address, err := common.HexToAddress(addressHex)
	if err != nil {
		panic(err)
	}

	return NewAddressLocation(address, name)
}

func NewAddressLocation(address common.Address, name string) common.AddressLocation {
	return common.AddressLocation{
		Address: address,
		Name:    name,
	}
}
