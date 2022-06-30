package main

import (
	"fmt"
	"strconv"
)

type Hanoi struct {
	name string
	data []int
}

func solveHanoiTower2() {
	var source = Hanoi{name: "source", data: []int{1, 2, 3, 4}}
	var target = Hanoi{name: "target"}
	var help1 = Hanoi{name: "help1"}
	var help2 = Hanoi{name: "help2"}
	solveHanoiTowerHelp2(&source, &help1, &help2, &target, len(source.data))
}

func solveHanoiTowerHelp2(source *Hanoi, help1 *Hanoi, help2 *Hanoi, target *Hanoi, sourceDataLength int) {
	if sourceDataLength == 1 {
		fmt.Print("moving " + strconv.Itoa(source.data[0]) + " from " + source.name + " to " + target.name)
		target.data = append([]int{source.data[0]}, target.data...)
		source.data = source.data[1:]
		// printAllStack(source, help, target)
		return
	}

	// solveHanoiTowerHelp(source, target, help, sourceDataLength-1)
	lastdata := source.data[0]
	fmt.Print("moving " + strconv.Itoa(lastdata) + " from " + source.name + " to " + target.name)
	target.data = append([]int{lastdata}, target.data...)
	source.data = source.data[1:]
	// printAllStack(source, help, target)
	// solveHanoiTowerHelp(help, source, target, sourceDataLength-1)
}

func solveHanoiTower() {
	var source = Hanoi{name: "source", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	var target = Hanoi{name: "target"}
	var help = Hanoi{name: "help"}
	solveHanoiTowerHelp(&source, &help, &target, len(source.data))
}

func solveHanoiTowerHelp(source *Hanoi, help *Hanoi, target *Hanoi, sourceDataLength int) {
	if sourceDataLength == 1 {
		fmt.Print("moving " + strconv.Itoa(source.data[0]) + " from " + source.name + " to " + target.name)
		target.data = append([]int{source.data[0]}, target.data...)
		source.data = source.data[1:]
		printAllStack(source, help, target)
		return
	}

	solveHanoiTowerHelp(source, target, help, sourceDataLength-1)
	lastdata := source.data[0]
	fmt.Print("moving " + strconv.Itoa(lastdata) + " from " + source.name + " to " + target.name)
	target.data = append([]int{lastdata}, target.data...)
	source.data = source.data[1:]
	printAllStack(source, help, target)
	solveHanoiTowerHelp(help, source, target, sourceDataLength-1)
}

func printAllStack(source *Hanoi, help *Hanoi, target *Hanoi) {
	fmt.Print(", Now the Stacks are:")
	var hanois = []*Hanoi{source, help, target}

	sourceHanoi := findByName(hanois, "source")
	printAllStackHelp(sourceHanoi)
	helpHanoi := findByName(hanois, "help")
	printAllStackHelp(helpHanoi)
	targetHanoi := findByName(hanois, "target")
	printAllStackHelp(targetHanoi)

	fmt.Println("")
}

func findByName(hanois []*Hanoi, targetName string) *Hanoi {
	for _, item := range hanois {
		if item.name == targetName {
			return item
		}
	}
	return nil
}

func printAllStackHelp(hanoi *Hanoi) {
	fmt.Print(" ")
	fmt.Print(hanoi.name)
	fmt.Print(":")
	for i, item := range hanoi.data {
		fmt.Print(item)
		if i < len(hanoi.data)-1 {
			fmt.Print(",")
		}
	}
}
