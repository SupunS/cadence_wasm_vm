(module

    (import "" "new_struct" (func $new_struct (result externref)))
    (import "" "new_struct_value" (func $new_struct_value (param externref) (param externref) (param i32) (result externref)))
    (import "" "string_const" (func $string_const (param i32) (param i32) (result externref)))
    (import "" "new_address_loaction" (func $new_address_loaction (param externref) (param externref) (result externref)))

    (memory (export "memory") 1)

    (data (i32.const 0) "0100000000000000")
    (data (i32.const 16) "TestContract")
    (data (i32.const 28) "Foo")

    (func $recursive_fib (param $n i32) (result i32)
        (if (i32.lt_s (local.get $n) (i32.const 2))
            (return (local.get $n))
        )
        (local.get $n)
        (i32.const 1)
        (i32.sub)
        (call $recursive_fib)
        (local.get $n)
        (i32.const 2)
        (i32.sub)
        (call $recursive_fib)
        (i32.add)
    )

    (func $imperative_fib (param $a i32) (result i32)
        (local $b i32) (local $c i32) (local $d i32) (local $e i32)

        (local.set $b (i32.const 1))
        (local.set $c (i32.const 1))
        (local.set $d (local.get $a))
        (local.set $e (i32.const 2))
        (block (loop
            (br_if 1 (i32.ge_s (local.get $e) (local.get $a)))
            (local.set $d (i32.add (local.get $b) (local.get $c)))
            (local.set $b (local.get $c))
            (local.set $c (local.get $d))
            (local.set $e (i32.add (local.get $e) (i32.const 1)))
            (br 0)
        ))
        (local.get $d)
    )

    (func $create_struct_simple (result externref)
        (call $new_struct)
    )

    (func $create_struct
        (local $a i32)
        (local.set $a (i32.const 0))
        (block (loop
            (br_if 1 (i32.ge_s (local.get $a) (i32.const 5)))
            (call $new_struct)
            drop
            (local.set $a (i32.add (local.get $a) (i32.const 1)))
            (br 0)
        ))
    )

    (func $empty_function)

    (func $create_composite_value (result externref)
        ;; Load address hex
        ;; string_const(index , len)
        (call $string_const (i32.const 0) (i32.const 16))

        ;; Load contract name
        ;; string_const(indetifier_index , identifier_len)
        (call $string_const (i32.const 16) (i32.const 12))

        ;; Create an address location
        (call $new_address_loaction)

        ;; Load struct name
        ;; string_const(indetifier_index , identifier_len)
        (call $string_const (i32.const 28) (i32.const 3))

        ;; Load composiote kind=1 (struct)
        (i32.const 1)

        ;; Create new comspoite value
        (call $new_struct_value)
    )

    (func $benchmark_composite_value
        (local $a i32)
        (local.set $a (i32.const 0))
        (block (loop
            (br_if 1 (i32.ge_s (local.get $a) (i32.const 5)))

            (call $create_composite_value)
            drop

            (local.set $a (i32.add (local.get $a) (i32.const 1)))
            (br 0)
        ))
    )

    (export "recursive_fib" (func $recursive_fib))
    (export "imperative_fib" (func $imperative_fib))
    (export "create_struct_simple" (func $create_struct_simple))
    (export "empty_function" (func $empty_function))
    (export "create_struct" (func $create_struct))
    (export "create_composite_value" (func $create_composite_value))
    (export "benchmark_composite_value" (func $benchmark_composite_value))
)
