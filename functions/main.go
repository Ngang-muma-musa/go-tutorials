package main

import (
	"fmt"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n/2; i++ {
		if n%2 == 0 {
			return false
		}
	}
	return true
}

func printPrimeNumbers(n int, withGoRoutine bool, prefix string) {
	t1 := time.Now()
	var wg sync.WaitGroup
	wg.Add(n)

	fmt.Print(prefix)
	for i := 0; i < n; i++ {
		var p bool
		_i := i
		f := func(v int) {
			p = isPrime(v)
			if p {
				//fmt.Print(_i, ", ")
			}
			wg.Done()
		}
		if withGoRoutine {
			go f(_i)
		} else {
			f(_i)
		}
	}
	wg.Wait()
	duration := time.Since(t1)
	fmt.Println("dur: ", duration)
}

func main() {
	printPrimeNumbers(1000000, true, "with go routine: ")
	printPrimeNumbers(1000000, false, "without go routine: ")
}
