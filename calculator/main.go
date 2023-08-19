package main

import (
	"fmt"
	"hash/maphash"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Start Questions")

	i := 1
	mark := 0
	for i <= 20 {
		i++
		if i > 0 && i <= 5 {
			generateQuestion(0, 0, &mark, i)
			continue
		} else if i > 5 && i <= 10 {
			generateQuestion(0, 10, &mark, i)
			continue
		} else if i > 10 && i <= 15 {
			generateQuestion(10, 10, &mark, i)
			continue
		} else if i > 15 && i <= 20 {
			continue
		} else {
			break
		}
	}
	wg.Wait()
	fmt.Println("Your score is ", mark, " / ", i)
}
func generateRandomNumbers(level int) int {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	return r.Intn(8) + level
}

func generateQuestion(num1Difficulty int, num2Difficulty int, mark *int, n int) {

	num1 := generateRandomNumbers(num1Difficulty)
	num2 := generateRandomNumbers(num2Difficulty)
	w := false
	c1 := make(chan bool)

	fmt.Println(num1, " X ", num2, " = ?")

	wg.Add(n)
	go func() {
		c1 <- false
		w = waitTimer()
		if w {
			fmt.Println("TIME OUT")
			c1 <- w
			return
		}
		wg.Done()
	}()

	for <-c1 == false {
		fmt.Println("in here")
		var userAns int
		fmt.Scanf("%d", &userAns)
		if num1*num2 == userAns {
			*mark++
			return
		}
	}
	return
}

func waitTimer() bool {
	newtimer := time.NewTimer(10 * time.Second)
	<-newtimer.C
	return true
}
