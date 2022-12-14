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
//	err := vm.LoadWasmFile("fib.wasm")
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
