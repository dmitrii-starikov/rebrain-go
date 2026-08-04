package main

import (
	"fmt"
	customNameForUnsafe "unsafe"
)

func anotherFile() {
	var s []int
	fmt.Println(len(s), cap(s))
	fmt.Printf("pntr: address=%p\n", &s)
	printSliceInfo(&s, "start")

	s = []int{4, 5, 6, 7}
	fmt.Println(len(s), cap(s))
	fmt.Printf("pntr: address=%p\n", &s)
	printSliceInfo(&s, "after 1 assign")

	oldP := &s
	oldData := make([]int, len(s))
	copy(oldData, s)

	s = []int{1, 2, 3, 4, 5}
	fmt.Println(len(s), cap(s))
	fmt.Printf("pntr: address=%p\n", &s) // address of slice doesn't change
	printSliceInfo(&s, "after 2 assign")

	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	fmt.Println(len(s), cap(s))
	fmt.Printf("pntr: address=%p\n", &s)
	printSliceInfo(&s, "after 3 assign")

	printSliceInfo(oldP, "old pointer after 1 assign") // changed

	printSliceInfo(&oldData, "old data") // changed

	////////////////////////////////////
	arr := [2]int{1, 2}

	fmt.Printf("%T\n", arr)
	fmt.Printf("%T\n", s)

	st := []string{"e", "l", "e", "m", "e", "n", "t"}
	st[0] = "E"
	fmt.Println(st)

	made := make([]int, 10, 100)
	fmt.Printf("pntr: address=%p\n", &made)
	fmt.Println(made, len(made), cap(made))
	fmt.Println()

	////////////////////////////////////
	s1 := []int{1, 2, 3}
	p1 := []int{4, 5, 6}

	printSliceInfo(&s1, "s1")
	s1 = append(s1, 8, 8)
	s1 = append(s1, p1...)
	fmt.Println(s1)
	printSliceInfo(&s1, "s1'")

	////////////////////////////////////
	s0 := []string{"e", "l", "e", "m", "e", "n", "t"}
	s11 := s0[:4]
	s21 := s0[1:6]
	fmt.Println(s11, s21) // [e l e m] [l e m e n]
	s11[1] = "L"          // slice `s11` - pointer on `s0` => `s0`, `s21` (also pointer) were changed!
	fmt.Println(s0, s21)  // [e L e m e n t] [L e m e n]

	fmt.Println(cap(s11)) // 7

	printStringSliceInfo(&s11, "S1 before append()")

	//s21[4] = "N"
	//s21[2] = "M" // still affect s11
	s11 = append(s11, "one", "two", "three", "four")
	fmt.Println(s0, s11, s21) // [e L e m e n t] [e L e m one two three four] [L e m e n]
	fmt.Println(cap(s11))     // 14
	s21[4] = "N"
	s21[2] = "M" // no affect s11
	// append() increased s11's capacity and allocated a new underlying array,
	// so modifications to s21 don't affect s11 anymore
	fmt.Println(s0, s11, s21) // [e L e M e N t] [e L e m one two three four] [L e M e N]

	printStringSliceInfo(&s11, "S1 after append()")

	//////////////////////////
	s8 := []int{1, 2, 3}
	p8 := []int{4, 5, 6}

	s8 = append(s8, p8...)
	fmt.Println(s8) //[1 2 3 4 5 6]

	s8 = append(s8, 22, 33, 44)
	fmt.Println(s8) //[1 2 3 4 5 6 22 33 44]

	////////////////////////

	stringSlice := make([]string, 10, 100)
	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice)) // [         ] 10 100
	stringSlice1 := stringSlice[:5]
	fmt.Println(stringSlice1, len(stringSlice1), cap(stringSlice1)) // [    ] 5 100
	stringSlice[0], stringSlice[1], stringSlice[2], stringSlice[3], stringSlice[4] = "A", "B", "C", "D", "E"
	fmt.Println(stringSlice1, len(stringSlice1), cap(stringSlice1)) // [A B C D E] 5 100
	stringSlice1 = append(stringSlice1, "F", "G", "H")
	fmt.Println(stringSlice1, len(stringSlice1), cap(stringSlice1)) // [A B C D E F G H] 8 100
	stringSlice[4] = "X"                                            // affect stringSlice1 because of unchanged capacity
	fmt.Println(stringSlice1, len(stringSlice1), cap(stringSlice1)) // [A B C D X F G H] 8 100

	fmt.Println(stringSlice, len(stringSlice), cap(stringSlice)) // [A B C D X F G H  ] 10 100 - also changed

	////////////////////////

	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3, 10)

	qnt := copy(dst, src)
	fmt.Println(qnt, dst) // 3 [1 2 3]

	src1 := "string"
	byteDst := make([]byte, 14, 18)
	qnt1 := copy(byteDst, src1)
	fmt.Println(qnt1, byteDst) // 6 [115 116 114 105 110 103 0 0 0 0 0 0 0 0]
}

func printSliceInfo(s *[]int, label string) {
	sliceHeader := (*[3]uintptr)(customNameForUnsafe.Pointer(s)) // struct {Pointer, len, cap}
	arrayPtr := customNameForUnsafe.Pointer(sliceHeader[0])

	fmt.Printf("%s:\n", label)
	fmt.Printf("  Slice structure address: %p\n", s)
	fmt.Printf("  Array structure address: %p\n", arrayPtr)
	fmt.Printf("  Length: %d, Capacity: %d\n", len(*s), cap(*s))
	fmt.Printf("  Struct: %v\n", *sliceHeader)
	fmt.Printf("  Data (address=%v): %v\n\n", arrayPtr, *s)
}

func printStringSliceInfo(s *[]string, label string) {
	sliceHeader := (*[3]uintptr)(customNameForUnsafe.Pointer(s)) // struct {Pointer, len, cap}
	arrayPtr := customNameForUnsafe.Pointer(sliceHeader[0])

	fmt.Printf("%s:\n", label)
	fmt.Printf("  Slice structure address: %p\n", s)
	fmt.Printf("  Array structure address: %p\n", arrayPtr)
	fmt.Printf("  Length: %d, Capacity: %d\n", len(*s), cap(*s))
	fmt.Printf("  Struct: %v\n", *sliceHeader)
	fmt.Printf("  Data (address=%v): %v\n\n", arrayPtr, *s)
}

func ExportedFunc(a int) int {
	return a + a
}
