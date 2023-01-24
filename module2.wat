(module

    ;; Built-in functions.
    ;; Module name for built-in functions is always empty. i.e: "".
    (import "" "string_const" (func $string_const (param i32) (param i32) (result externref)))
    (import "" "new_composite_value" (func $new_composite_value (param externref) (param externref) (param i32) (result externref)))
    (import "" "new_address_location" (func $new_address_location (param externref) (param externref) (result externref)))
    (import "" "set_int_member" (func $set_int_member (param $self externref) (param $fieldName externref) (param i32)))
    (import "" "set_member" (func $set_member (param $self externref) (param $fieldName externref) (param externref)))

    (memory (export "memory") 1)

    (data (i32.const 0) "0100000000000000")
    (data (i32.const 16) "ImportedContract")
    (data (i32.const 28) "Foo")
    (data (i32.const 31) "a")
    (data (i32.const 32) "b")
    (data (i32.const 33) "Hello")

    ;; Suppose the 'Foo' constructor is:
    ;; init () {
    ;;   self.a = 4
    ;;   self.b = "Hello"
    ;; }
    (func (export "Foo") (result externref)
        (local $self externref)

        ;; --- Go composite value creation ---

        ;; Load address hex
        ;; string_const(index , len)
        (call $string_const (i32.const 0) (i32.const 16))

        ;; Load contract name
        ;; string_const(indetifier_index , identifier_len)
        (call $string_const (i32.const 16) (i32.const 12))

        ;; Create an address location
        (call $new_address_location)

        ;; Load struct name
        ;; string_const(indetifier_index , identifier_len)
        (call $string_const (i32.const 28) (i32.const 3))

        ;; Load composiote kind=1 (struct)
        (i32.const 1)

        ;; Create new comspoite value
        call $new_composite_value
        local.set $self

        ;; --- Constructor body ---

        ;; Load self
        local.get $self
        ;; Load filedName
        (call $string_const (i32.const 31) (i32.const 1))
        ;; load int value
        i32.const 4
        ;; Set 'a'
        call $set_int_member

        ;; Load self
        local.get $self
        ;; Load filedName
        (call $string_const (i32.const 32) (i32.const 1))
        ;; Load 'Hello'
        (call $string_const (i32.const 33) (i32.const 5))
        ;; Set 'b'
        call $set_member

        ;; Return self
        local.get $self
    )
)
