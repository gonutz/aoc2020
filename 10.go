//+build ignore

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var ns []int
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(line)
		ns = append(ns, n)
	}
	sort.Ints(ns)
	ns = append([]int{0}, ns...)
	ns = append(ns, ns[len(ns)-1]+3)
	var ones, threes int
	for i := 1; i < len(ns); i++ {
		d := ns[i] - ns[i-1]
		if d == 1 {
			ones++
		}
		if d == 3 {
			threes++
		}
	}
	fmt.Println(ones, "*", threes, "=", ones*threes)

	ways := countWays(ns, make(map[*int]int64))
	fmt.Println(ways, "ways")
}

func countWays(x []int, cache map[*int]int64) int64 {
	// This function started out as tree-nesting recursion. Like calculating
	// Fibonacci numbers by recursion, subsequent numbers are calculated over
	// and over again. To solve this you would usually remember some state
	// variables by making the function iterative and using a for-loop.
	// I chose Mike Acton's way
	// 	https://www.youtube.com/watch?v=x61H6qEtK08
	// and just cache results in a map instead, keeping the function recursive.
	// Since we cannot use []int as a key in a map, we use the pointer to the
	// first item to identify that we have been here before.
	if n, ok := cache[&x[0]]; ok {
		return n
	}
	var sum int64
	if len(x) == 1 {
		sum = 1
	} else {
		if len(x) >= 2 && x[1]-x[0] <= 3 {
			sum += countWays(x[1:], cache)
		}
		if len(x) >= 3 && x[2]-x[0] <= 3 {
			sum += countWays(x[2:], cache)
		}
		if len(x) >= 4 && x[3]-x[0] <= 3 {
			sum += countWays(x[3:], cache)
		}
	}
	cache[&x[0]] = sum
	return sum
}

const input = `18
47
144
147
124
45
81
56
16
59
97
83
75
150
33
165
30
159
84
141
104
25
164
90
92
88
2
8
51
24
153
63
27
123
127
58
108
52
38
15
149
66
72
21
46
89
135
55
34
37
78
65
134
148
76
138
103
162
114
109
42
77
102
163
7
105
69
39
91
111
131
130
6
137
96
82
64
3
95
136
85
9
116
17
99
12
117
62
50
110
26
115
71
57
156
120
98
1
70`
