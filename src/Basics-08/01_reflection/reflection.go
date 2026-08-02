package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string `json:"name"`
	Age      uint   `json:"age"`
}

func main() {
	var v interface{} = "value"
	reflectValue := reflect.ValueOf(v)
	valueType := reflectValue.Type()
	typeName := valueType.Name() // string

	fmt.Printf("Type: %s\n", typeName)

	user := User{"ivan", 25}
	reflectTypeUser := reflect.TypeOf(user)
	field := reflectTypeUser.Field(0)
	fmt.Println(field.Name, field.Type.Name(), field.Tag)

	var x float64 = 2.9
	vx := reflect.ValueOf(x)
	//vx.SetFloat(0.8) // Error
	// panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Printf("CanSet vx: %v\n", vx.CanSet()) // false

	var x2 float64 = 2.9
	vx2 := reflect.ValueOf(&x2)
	e := vx2.Elem()
	e.CanSet() // true
	e.SetFloat(1.1)
	fmt.Printf("x2 = %.3f\n", x2) // 1.100

	type MyInt int
	var xy MyInt = 12
	vy := reflect.ValueOf(xy)
	fmt.Println("type:", vy.Type().Name())                   // MyInt.
	fmt.Println("kind is uint8: ", vy.Kind() == reflect.Int) // true.

	type Data struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}
	var xy2 Data
	t := reflect.TypeOf(xy2)
	fmt.Println("json tag of first field:", t.Field(0).Tag.Get("json"))  // foo.
	fmt.Println("json tag of second field:", t.Field(1).Tag.Get("json")) // bar.
}
