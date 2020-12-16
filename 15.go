//+build ignore

package main

import "fmt"

func main() {
	input := []int{2, 0, 1, 7, 4, 14, 18}
	m := make(map[int]int)
	for i, n := range input[:len(input)-1] {
		m[n] = i
	}
	next := input[len(input)-1]
	for i := len(input); i < 30000000; i++ {
		index, ok := m[next]
		m[next] = i - 1
		if ok {
			next = i - index - 1
		} else {
			next = 0
		}
		if i == 2020-1 {
			fmt.Println(next)
		}
	}
	fmt.Println(next)
}
