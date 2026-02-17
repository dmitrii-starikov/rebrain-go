package main

import (
	"fmt"
	"time"
)

func pingPong(name string, hits chan int) {
	for {
		hit := <-hits // waiting for partner hit
		fmt.Println(name, "hit", hit)
		time.Sleep(time.Millisecond * 500)
		hits <- hit + 1 // our hit
	}
}

func main() {
	table := make(chan int) // unbuffered

	go pingPong("Player1", table)
	//go pingPong("Player2", table)
	//go pingPong("Player3", table)

	table <- 1 // start the game
	time.Sleep(time.Second * 5)
}
