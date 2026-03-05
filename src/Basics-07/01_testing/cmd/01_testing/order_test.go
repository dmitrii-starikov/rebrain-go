package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOrder(t *testing.T) {
	req := require.New(t)
	orders, err := GetOrder(5)
	req.NoError(err)
	req.Len(orders, 5)
	req.NotEqual(1015, orders[0].ID)
}
