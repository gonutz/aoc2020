//+build ignore

package main

import "fmt"

func main() {
	const x = 0
	input := []int64{23, x, x, x, x, x, x, x, x, x, x, x, x, 41, x, x, x, 37, x, x, x, x, x, 421, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, x, 17, x, 19, x, x, x, x, x, x, x, x, x, 29, x, 487, x, x, x, x, x, x, x, x, x, x, x, x, 13}
	stamp := int64(1002392)
	fmt.Println("earliest", earliest(stamp, input))
	fmt.Println("timestamp", findTime(input))
}

func earliest(stamp int64, ns []int64) int64 {
	min := ns[0]
	minD := int64(99999999999999)
	for _, n := range ns {
		if n == 0 {
			continue
		}
		d := (n - stamp%n) % n
		if d < minD {
			min = n
			minD = d
		}
	}
	return min * minD
}

func findTime(stamps []int64) int64 {
	type constraint struct {
		mod    int64
		yields int64
	}
	var constraints []constraint
	for i, n := range stamps {
		if n == 0 {
			continue
		}
		constraints = append(constraints, constraint{
			mod:    n,
			yields: (n + (n-int64(i))%n) % n,
		})
	}
	var n, inc int64 = 0, 1
	for len(constraints) > 0 {
		for n%constraints[0].mod != constraints[0].yields {
			n += inc
		}
		inc *= constraints[0].mod
		constraints = constraints[1:]
	}
	return n
}
