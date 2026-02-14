package main

import (
	"context"
	"fmt"
)

// define our own type for the key to avoid collisions.
type ctxKey string

const userIDKey ctxKey = "userID"

func main() {
	// create a root context and add userID
	ctx := context.WithValue(context.Background(), userIDKey, "12345")

	// передаём контекст в функцию, где он будет извлечён
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	// trying to read the value by the userIDKey key
	if uid, ok := ctx.Value(userIDKey).(string); ok {
		fmt.Println("User ID:", uid)
	} else {
		fmt.Println("User ID not found in context")
	}
}
