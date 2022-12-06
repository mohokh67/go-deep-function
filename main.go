package main

import (
	"errors"
	"fmt"
	"go-deep-function/visitor"
	"math"
)

func main() {
	// fmt.Println("There are tons of things you can do with functions in GO!")
	vi := visitor.New()
	vi.Join()
	vi.Join()
	vi.Join()
	vi.Join()
	vi.Left()
	println(vi.ActiveString()) // 3
	println(vi.TotalString())  // 4

	bulkJoin := vi.BulkActivity(string(visitor.AddActivity))
	bulkLeft := vi.BulkActivity(string(visitor.SubtractActivity))

	bulkJoin(2)
	println(vi.ActiveString()) // 13
	println(vi.TotalString())  // 14

	bulkLeft(5)
	println(vi.ActiveString()) // 8
	println(vi.TotalString())  // 14

	// anonymous := func() {
	// 	fmt.Println("anonymous function")
	// }
	// anonymous()
}

// stateful func
func powerOfTwo() func() int {
	x := 1.0
	return func() int {
		x += 1
		return int(math.Pow(x, 2))
	}
}

// Variadic parameters
func sum(values ...int) (total int) {
	for _, value := range values {
		total += value
	}
	return
}

func greeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can't be empty")
	}
	return fmt.Sprintf("Hi %s, Let's do some golang func practice!\n", name), nil
}

func mustGreet(val string, err error) string {
	if err != nil {
		return fmt.Sprintf("error: %s\n", err)
	}
	return val
}
