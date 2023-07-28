package main

import "fmt"

func main() {

	fmt.Println("\nPrint multiples of 3 between 1 to 100")
	i := 1
	for i < 100 {
		if i%3 == 0 {
			fmt.Println(i)
		}
		i++
	}

	fmt.Println("\nPrint Feez for multiples of 3 and Buzz for multiples of 5, FeezBuzz if a number is both")

	j := 1
	for j <= 100 {
		if j%3 == 0 && j%5 == 0 {
			fmt.Println("FeezBuzz")
		} else if j%5 == 0 {
			fmt.Println("Buzz")
		} else if j%3 == 0 {
			fmt.Println("Feez")
		} else {
			fmt.Println(j)
		}
		j++
	}
}

/*
	1) small
*/
