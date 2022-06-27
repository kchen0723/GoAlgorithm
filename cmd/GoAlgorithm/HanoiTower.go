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
	var source = Hanoi{name: "source", data: []int{1, 2, 3, 4}}
	var target = Hanoi{name: "target"}
	var help = Hanoi{name: "help"}
	solveHanoiTowerHelp(&source, &help, &target)
}

func solveHanoiTowerHelp(source *Hanoi, help *Hanoi, target *Hanoi) {
	if len(source.data) == 1 {
		fmt.Println("moving " + strconv.Itoa(source.data[0]) + " from: " + source.name + " to: " + target.name)
		target.data = append([]int{source.data[0]}, target.data...)
		source.data = []int{}
		return
	}

	lastdata := source.data[len(source.data)-1]
	copieddata := make([]int, len(source.data)-1)
	copy(copieddata, source.data[:len(source.data)-1])
	source.data = copieddata
	solveHanoiTowerHelp(source, target, help)
	target.data = append([]int{lastdata}, target.data...)
	fmt.Println("moving " + strconv.Itoa(lastdata) + " from: " + source.name + " to: " + target.name)
	solveHanoiTowerHelp(help, source, target)
}
