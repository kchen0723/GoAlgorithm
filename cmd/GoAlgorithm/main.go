package main

import (
	"fmt"
	"math/rand"

	// "os"
	// "os/signal"
	// "syscall"
	"time"
)

var firstchan chan int

func main() {
	numbers := []int{1, 2, 3, 4}
	permutations := GetPermutationUnique(numbers, 3)
	permutations = GetPermutationMultipleTimes(numbers, 3)
	permutations = GetCombinationUnique(numbers, 3)
	permutations = GetCombinationMultipleTimes(numbers, 3)
	numbers = []int{2, 1, 2}
	permutations = GetPermutationDuplicate(numbers, 3)
	permutations = GetCombinationDuplicate(numbers, 2)
	if len(permutations) > 0 {
	}

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	Combinations := GetCombinationMultipleTimes(numbers, 4)
	var result [][]int
	for _, item := range Combinations {
		if Is4NumberCombationValid(item) {
			result = append(result, item)
		}
	}
	if len(result) > 0 {

	}
	// operatorsPermutation := get4NumberPostfixBackTrack([]int{8, 3, 8, 3})
	// if len(operatorsPermutation) > 0 {

	// }

	// tree := createCalculateTree()
	// treeString := sealize(tree)
	// fmt.Println(treeString)
	// rebuildTree := desealize(treeString)
	// if rebuildTree != nil {
	// }

	// input := "1+5-2*(4+3)"
	// tokens := GetTokens(input)
	// if len(tokens) > 0 {
	// 	postfixTokens := ParseExpression(tokens)
	// 	if len(postfixTokens) > 0 {
	// 		result := calculate(postfixTokens)
	// 		if result == 2 {
	// 			fmt.Println("bingo")
	// 		}
	// 	}
	// }
}

// func main() {
// 	firstchan = make(chan int, 10)
// 	// go pushToChan()
// 	// go takeWithTimeout()

// 	left := []int{1, 2, 3}
// 	right := []int{1, 3, 2}
// 	result := buildTree(left, right)
// 	if result != nil {

// 	}

// 	fmt.Printf("Press CTRL-C to stop\r\n")
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
// 	<-c
// 	fmt.Printf("Stopping on CTRL-C")
// }

func pushToChan() {
	// pushTicker := time.NewTicker(time.Second * 2)
	timeout := time.NewTimer(time.Second * 1)
	for i := 10; i < 100; i++ {
		fmt.Println(i)
		rand.Seed(time.Now().UnixNano())
		// left := rand.Intn(i)
		// right := left * left
		// c := add(left, right)
		select {
		// case firstchan <- c:
		// 	fmt.Println("push to chan: " + strconv.Itoa(c))
		// 	time.Sleep(time.Second * 2)
		// case <-time.After(time.Second * 1):
		// case <-pushTicker.C:
		case <-timeout.C:
			fmt.Println("cannot push to chan")
			time.Sleep(time.Second * 2)
		default:
			fmt.Println("default case")
			time.Sleep(time.Microsecond * 500)
		}
	}
	fmt.Println("finished push to chan")
}

func add(a int, b int) int {
	return a + b
}

func takeFromChan() {
	for {
		number := <-firstchan
		fmt.Println(number)
	}
}

func takeWithTimeout() {
	timeout := time.NewTimer(time.Second * 1)
	for {
		select {
		case number := <-firstchan:
			fmt.Println(number)
		// case <-time.After(time.Second * 1):
		case <-timeout.C:
			fmt.Println("timeout case")
		}
	}
}
