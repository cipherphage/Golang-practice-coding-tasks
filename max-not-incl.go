package main

import (
	"fmt"
)

func maxNotIncl(a []int) int {
	r := 0
	max := -1000000
	m := make(map[int]int)

	// add to map and also find max value
	for i := range a {
		m[a[i]] = a[i]
		if a[i] > max {
			max = a[i]
		}
	}

	// if max is zero or less, return 1
	if max <= 0 {
		r = 1
		return r
	}

	// check for missing nums
	for i := 1; i <= max; i++ {
		_, prs := m[i]

		// if num missing, hold only the latest
		if prs == false {
			r = i
		}

		// if none are missing, return max + 1
		if i == max && r == 0 {
			r = max + 1
		}
	}

	return r
}

func main() {
	fmt.Println(maxNotIncl([]int{9, 3, 6, 4, 1, 2, 7, 1})) // 8
	fmt.Println(maxNotIncl([]int{1, 2, 3}))                // 4
	fmt.Println(maxNotIncl([]int{-1, -3}))                 // 1
	fmt.Println(maxNotIncl([]int{0}))                      // 1
	fmt.Println(maxNotIncl([]int{-1}))                     // 1
	fmt.Println(maxNotIncl([]int{1}))                      // 2
	fmt.Println(maxNotIncl([]int{10, 10, 10}))             // 9
}
