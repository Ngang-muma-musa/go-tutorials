package main

import "fmt"

func main() {
	x := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17, 3, 2,
	}
	min := x[0]

	for i := range x {
		if x[i] < min {
			min = x[i]
		}
	}

	fmt.Print("The smallest element in the array is: ", min)
}

/*
1) In Go array indexing starts from 0 so to access the forth element will be arr[3]
2) The length of the slice make([]int, 3, 9) is 3 with capacity 9
3) given the following array x := [6]string{"a","b","c","d","e","f"}, x[2:5] gives you [c,d,e]
*/
