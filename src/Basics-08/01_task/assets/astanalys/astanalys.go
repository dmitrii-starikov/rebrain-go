package astanalys

import (
	"context"
	"errors"
	"fmt"
	"log"
)

const (
	one = iota
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
)

var (
	firstName = "Ivan"
	lastName  = "Ivanov"
	phone     = "89999999999"
	city      = "moscow"
)

func exampleFunc(ctx context.Context) error {
	var o string

	a := firstName
	b := a
	a = lastName
	c := phone

	for i := 0; i < 100; i++ {
		j := i
		log.Println(j)
	}

	o = city
	fmt.Println(a, b, c, o)

	if a != firstName {
		return errors.New("error")
	}

	return nil
}
