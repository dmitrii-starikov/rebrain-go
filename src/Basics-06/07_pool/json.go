package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 1024)
	},
}

func marshalUser(user User) ([]byte, error) {
	buf := bufferPool.Get().([]byte)
	buf = buf[:0] // clear length, but save capacity

	result, err := json.Marshal(user)
	if err != nil {
		bufferPool.Put(buf)
		return nil, err
	}

	buf = append(buf, result...)

	defer bufferPool.Put(buf)

	return result, nil
}

func main() {
	user := User{ID: 1, Name: "John", Email: "john@example.com"}

	data, _ := marshalUser(user)
	fmt.Println(string(data))
}
