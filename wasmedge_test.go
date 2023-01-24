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

//
//func TestRecursiveFib(t *testing.T) {
//
//	wasmedge.SetLogErrorLevel()
//
//	var conf = wasmedge.NewConfigure(wasmedge.WASI)
//	var vm = wasmedge.NewVMWithConfig(conf)
//	var wasi = vm.GetImportModule(wasmedge.WASI)
//	wasi.InitWasi(
//		os.Args[1:],     // The args
//		os.Environ(),    // The envs
//		[]string{".:."}, // The mapping preopens
//	)
//
//	/// Instantiate wasm
//	err := vm.LoadWasmFile("module1.wasm")
//	require.NoError(t, err)
//
//	err = vm.Validate()
//	require.NoError(t, err)
//
//	bg := bindgen.New(vm)
//	bg.Instantiate()
//
//	// Gets the `recursive_fib` exported function from the WebAssembly instance.
//	res, _, err := bg.Execute("recursive_fib", "14")
//
//	assert.Equal(t, int32(377), res)
//}
