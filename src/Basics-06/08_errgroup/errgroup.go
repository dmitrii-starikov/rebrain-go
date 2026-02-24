package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func runTasks() error {
	g := new(errgroup.Group)

	tasks := []func() error{
		func() error { time.Sleep(2 * time.Second); return errors.New("error in task 1") },
		func() error { time.Sleep(1 * time.Second); return nil },
		func() error { time.Sleep(3 * time.Second); return nil },
	}

	for _, task := range tasks {
		task := task // capture to local scope
		g.Go(func() error {
			return task()
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := runTasks(); err != nil {
		fmt.Println("runTasks:", err)
	} else {
		fmt.Println("ok")
	}
}
