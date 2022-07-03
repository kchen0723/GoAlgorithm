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
	var help = Hanoi{name: "help"}
	var target1 = Hanoi{name: "target1"}
	var target2 = Hanoi{name: "target2"}

	solveHanoiTowerHelp2(&source, &help, &target1, &target2, len(source.data))
}

func solveHanoiTowerHelp2(source *Hanoi, help *Hanoi, target1 *Hanoi, target2 *Hanoi, sourceDataLength int) {
	if sourceDataLength == 1 {
		fmt.Println("moving " + strconv.Itoa(source.data[0]) + " from " + source.name + " to " + target1.name)
		target1.data = append([]int{source.data[0]}, target1.data...)
		source.data = source.data[1:]
		// printAllStack(source, help, target)
		return
	}

	if sourceDataLength == 2 {
		fmt.Println("moving " + strconv.Itoa(source.data[0]) + " from " + source.name + " to " + target1.name)
		target1.data = append([]int{source.data[0]}, target1.data...)
		source.data = source.data[1:]
		if len(source.data) >= 1 {
			fmt.Println("moving " + strconv.Itoa(source.data[0]) + " from " + source.name + " to " + target2.name)
			target2.data = append([]int{source.data[0]}, target2.data...)
			source.data = source.data[1:]
		}
		// printAllStack(source, help, target)
		return
	}

	solveHanoiTowerHelp2(source, target1, target2, help, sourceDataLength-1)
	lastdata := source.data[0]
	fmt.Print("moving " + strconv.Itoa(lastdata) + " from " + source.name + " to " + target1.name)
	target1.data = append([]int{lastdata}, target1.data...)
	source.data = source.data[1:]
	// printAllStack(source, help, target)
	solveHanoiTowerHelp2(help, source, target1, target2, sourceDataLength-2)
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
