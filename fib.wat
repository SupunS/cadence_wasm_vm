(module

    (import "" "new_struct" (func $new_struct (result externref)))

    (memory 1)


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
    (export "recursive_fib" (func $recursive_fib))

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
    (export "imperative_fib" (func $imperative_fib))

    (func $create_struct_simple (result externref)
        (call $new_struct)
    )
    (export "create_struct_simple" (func $create_struct_simple))

    (func $create_struct
        (local $a i32)
        (local.set $a (i32.const 0))
        (block (loop
            (br_if 1 (i32.ge_s (local.get $a) (i32.const 7)))
            (call $new_struct)
            drop
            (local.set $a (i32.add (local.get $a) (i32.const 1)))
            (br 0)
        ))
    )
    (export "create_struct" (func $create_struct))

    (func $empty_function)
    (export "empty_function" (func $empty_function))
)
