package main

import (
	"fmt"
	"strconv"
)

type Hanoi struct {
	name string
	data []int
}

func solveHanoiTower() {
	var source = Hanoi{name: "source", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	var target = Hanoi{name: "target"}
	var help = Hanoi{name: "help"}
	solveHanoiTowerHelp(&source, &help, &target, len(source.data))
}

func solveHanoiTowerHelp(source *Hanoi, help *Hanoi, target *Hanoi, sourceDataLength int) {
	if sourceDataLength == 1 {
		fmt.Println("moving " + strconv.Itoa(source.data[0]) + " from: " + source.name + " to: " + target.name)
		target.data = append([]int{source.data[0]}, target.data...)
		source.data = source.data[1:]
		return
	}

	solveHanoiTowerHelp(source, target, help, sourceDataLength-1)
	lastdata := source.data[0]
	fmt.Println("moving " + strconv.Itoa(lastdata) + " from: " + source.name + " to: " + target.name)
	target.data = append([]int{lastdata}, target.data...)
	source.data = source.data[1:]
	solveHanoiTowerHelp(help, source, target, sourceDataLength-1)
}
