package main

import "fmt"
import "01testing/pkg/summy"

func main() {
	fmt.Println("Hello World")
	fmt.Printf("sum %d", summy.Sum(1, 7))
}

type Order struct {
	ID int
}

func GetOrder(num int) ([]*Order, error) {
	//Fail
	orders := []*Order{
		{ID: 123},
		{ID: 456},
		{ID: 789},
	}
	//Passed
	orders = []*Order{
		{ID: 1},
		{ID: 88},
		{ID: 123},
		{ID: 456},
		{ID: 789},
	}

	return orders, nil
}
