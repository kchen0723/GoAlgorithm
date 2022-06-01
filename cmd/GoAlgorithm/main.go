package main

import (
	"fmt"
	"math/rand"
	"regexp"

	// "os"
	// "os/signal"
	// "syscall"
	"time"
)

var firstchan chan int

func main() {
	// tree := createCalculateTree()
	// treeString := sealize(tree)
	// fmt.Println(treeString)

	input := "(-5+9)*(7+3)"
	re := regexp.MustCompile("([0-9\\.]+|[\\(\\)|+|\\-|*|/])")
	splitor := re.FindAllStringSubmatch(input, -1)
	for _, item := range splitor {
		fmt.Println(item[0])
	}
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
