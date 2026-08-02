## Tests

### Before
```
go@linux:~/GolandProjects/rebrain-go/src/Basics-08/01_task$ make test_01
=== RUN   TestStructToMap
=== RUN   TestStructToMap/simple_test
=== NAME  TestStructToMap
convertor_test.go:44:
Error Trace:    convertor_test.go:44
Error:          Expected value not to be nil.
Test:           TestStructToMap
=== NAME  TestStructToMap/simple_test
testing.go:1811: test executed panic(nil) or runtime.Goexit: subtest may have called FailNow on a parent test
--- FAIL: TestStructToMap (0.00s)
--- FAIL: TestStructToMap/simple_test (0.00s)
=== RUN   TestMapToStruct
=== RUN   TestMapToStruct/simple_test
=== NAME  TestMapToStruct
convertor_test.go:94:
Error Trace:    convertor_test.go:94
Error:          Not equal:
expected: "Superman"
actual  : ""

                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -Superman
                                +
                Test:           TestMapToStruct
=== NAME  TestMapToStruct/simple_test
testing.go:1811: test executed panic(nil) or runtime.Goexit: subtest may have called FailNow on a parent test
--- FAIL: TestMapToStruct (0.00s)
--- FAIL: TestMapToStruct/simple_test (0.00s)
=== RUN   TestStructToMapAndBack
=== RUN   TestStructToMapAndBack/simple_test
=== NAME  TestStructToMapAndBack
convertor_test.go:138:
Error Trace:    convertor_test.go:138
Error:          Expected value not to be nil.
Test:           TestStructToMapAndBack
=== NAME  TestStructToMapAndBack/simple_test
testing.go:1811: test executed panic(nil) or runtime.Goexit: subtest may have called FailNow on a parent test
--- FAIL: TestStructToMapAndBack (0.00s)
--- FAIL: TestStructToMapAndBack/simple_test (0.00s)
FAIL
FAIL    module07/internal/convertor     0.005s
FAIL
make: *** [Makefile:4: test_01] Error 1
```

Check [convertor_test.go](../internal/convertor/convertor_test.go)

### `TestStructToMap`

Tests converting struct to map:

-   simple test — basic field conversion
-   with struct field — nested structs remain as structs (not converted to maps)
-   more nesting — deep nested structures
-   not struct — non-struct input returns `nil`

### `TestMapToStruct`

Tests filling struct from map:
-   simple test — basic field mapping
-   with struct field — supports nested structs (values can be already-built structs)
-   more nesting — recursive struct filling

### `TestStructToMapAndBack`

Tests round-trip conversion: struct → map → struct:
-   simple test — basic round-trip
-   with pointer — pointer handling
-   more nesting — deep nesting

---

### After

```
=== RUN   TestStructToMap
=== RUN   TestStructToMap/simple_test
=== RUN   TestStructToMap/with_struct_field
=== RUN   TestStructToMap/more_nesting
=== RUN   TestStructToMap/not_struct
--- PASS: TestStructToMap (0.00s)
    --- PASS: TestStructToMap/simple_test (0.00s)
    --- PASS: TestStructToMap/with_struct_field (0.00s)
    --- PASS: TestStructToMap/more_nesting (0.00s)
    --- PASS: TestStructToMap/not_struct (0.00s)
=== RUN   TestMapToStruct
=== RUN   TestMapToStruct/simple_test
=== RUN   TestMapToStruct/with_struct_field
=== RUN   TestMapToStruct/more_nesting
--- PASS: TestMapToStruct (0.00s)
    --- PASS: TestMapToStruct/simple_test (0.00s)
    --- PASS: TestMapToStruct/with_struct_field (0.00s)
    --- PASS: TestMapToStruct/more_nesting (0.00s)
=== RUN   TestStructToMapAndBack
=== RUN   TestStructToMapAndBack/simple_test
=== RUN   TestStructToMapAndBack/with_pointer
=== RUN   TestStructToMapAndBack/more_nesting
--- PASS: TestStructToMapAndBack (0.00s)
    --- PASS: TestStructToMapAndBack/simple_test (0.00s)
    --- PASS: TestStructToMapAndBack/with_pointer (0.00s)
    --- PASS: TestStructToMapAndBack/more_nesting (0.00s)
PASS
ok  	module07/internal/convertor	0.004s
```

---

## Useful Reflection Functions in Go — Cheat Sheet

### 1. Basic Operations
```go
// Get reflect.Value
v := reflect.ValueOf(interface{})
// Get reflect.Type
t := reflect.TypeOf(interface{})
// Get value kind
kind := v.Kind() // reflect.Struct, reflect.String, reflect.Ptr, ...
// Check for nil
isNil := v.IsNil()
// Check if valid
isValid := v.IsValid()
```

### 2. Working with Structs

```go
// Number of fields
count := v.NumField()
// Get field by index
field := v.Field(i)
// Get field by name
field := v.FieldByName("Name")
// Get field type
fieldType := t.Field(i)
// Get field name
name := fieldType.Name
// Get field tag
tag := fieldType.Tag.Get("keyname")
// Check if field is exported
isExported := field.IsExported()
// Check if field can be set
canSet := field.CanSet()
```

### 3. Working with Pointers

```go
// Get underlying value from pointer
elem := v.Elem()
// Create new pointer
ptr := reflect.New(type)
// Check if it's a pointer
isPtr := v.Kind() == reflect.Ptr
```

### 4. Working with Maps

```go
// Create map
m := reflect.MakeMap(type)
// Set key-value pair
m.SetMapIndex(key, value)
// Get value by key
val := m.MapIndex(key)
// Iterate over map
iter := m.MapRange()
for iter.Next() {
key := iter.Key()
value := iter.Value()
}
```

### 5. Setting Values

```go
// Set value
field.Set(value)                    // interface{}
field.SetString("hello")            // string
field.SetInt(42)                    // int
field.SetUint(42)                   // uint
field.SetFloat(3.14)                // float
field.SetBool(true)                 // bool
field.SetBytes([]byte{1,2,3})       // []byte
// Set zero value
field.Set(reflect.Zero(field.Type()))
// Create new value
newVal := reflect.New(field.Type()).Elem()
```

### 6. Creating Values

```go
// Create new value of type
val := reflect.New(type).Elem()
// Create new pointer
ptr := reflect.New(type)
// Zero value
zero := reflect.Zero(type)
// Convert to interface{}
interfaceVal := val.Interface()
```

### 7. Type checking

```go
// Compare types
if val.Type() == targetType
// Check assignability
if val.Type().AssignableTo(targetType)
// Check convertibility
if val.Type().ConvertibleTo(targetType)
// Get underlying type (for pointers, slices, maps)
baseType := val.Type().Elem()
```

### 8. Working with Slices

```go
// Slice length
length := v.Len()
// Get element
elem := v.Index(i)
// Create slice
slice := reflect.MakeSlice(type, len, cap)
// Append element
slice = reflect.Append(slice, elem)
// Append slice
slice = reflect.AppendSlice(slice, otherSlice)
```

### 9. Calling methods

```go
// Get method by name
method := v.MethodByName("MethodName")
// Check if method exists
if method.IsValid() {
// Call method with arguments
result := method.Call([]reflect.Value{arg1, arg2})
}
// Call without arguments
result := method.Call(nil)
```

### 10. Working with Interfaces

```go
// Empty interface
type Empty interface{}
// Check if type implements interface
if v.Type().Implements(interfaceType) {
}
```