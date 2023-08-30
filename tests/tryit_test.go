package tests

import (
	"fmt"
	"testing"
)

func Test_Map(t *testing.T) {
	tData := []int{1, 2, 3, 4, 5}
	ret := Map(tData, func(index int, item int) int {
		return item + 1
	})
	fmt.Println(ret)
}

// Map returns a new slice with transformed elements
func Map[T any](items []T, fn func(index int, item T) T) []T {
	var mappedItems []T
	for index, value := range items {
		mappedItems = append(mappedItems, fn(index, value))
	}
	return mappedItems
}
