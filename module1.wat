(module

    ;; Built-in functions.
    ;; Module name for built-in functions is always empty. i.e: "".
    (import "" "new_struct" (func $new_struct (result externref)))
    (import "" "string_const" (func $string_const (param i32) (param i32) (result externref)))

    ;; Imported functions from other modules
    (import "0x2.ImportedContract" "Foo" (func $0x2.ImportedContract.Foo (result externref)))

    (memory (export "memory") 1)

    (func $create_composite_value (result externref)
        ;; Call 0x2.ImportedContract.Foo()
        call $0x2.ImportedContract.Foo
    )

    (func $benchmark_composite_value (param $count i32)
        (local $a i32)
        (local.set $a (i32.const 0))
        (block (loop
            (br_if 1 (i32.ge_s (local.get $a) (local.get $count)))

            (call $create_composite_value)
            drop

            (local.set $a (i32.add (local.get $a) (i32.const 1)))
            (br 0)
        ))
    )

    (export "create_composite_value" (func $create_composite_value))
    (export "benchmark_composite_value" (func $benchmark_composite_value))
)
