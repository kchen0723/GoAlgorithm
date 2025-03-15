package main

import "fmt"

func IterationArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func RecursiveArray(arr []int) {
	RecursiveArrayHelp(arr, 0)
}

func RecursiveArrayHelp(arr []int, index int) {
	if index >= 0 && index < len(arr) {
		fmt.Println(arr[index])
		RecursiveArrayHelp(arr, index+1)
	}
}
