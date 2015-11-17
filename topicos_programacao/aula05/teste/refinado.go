package main

import (
	"fmt"
	"time"
)

var returnList = make([][]int, 0)

func containValueOnList(list []int, x int) (int, bool) {
	for index, value := range list {
		if value == x {
			return index, true
		}
	}
	return 0, false
}

func testState(list []int, size int) bool {
	if len(list) == size {
		returnList = append(returnList, list)
		return false
	}
	if len(list) > size {
		return false
	}
	return true
}

func calcNextSteps(n int, list []int) []int {
	tmp := make([]int, 0)
	start := 1
	if len(list) != 0 {
		start = list[len(list)-1] + 1
	}
	for i := start; i <= n; i++ {
		tmp = append(tmp, i)
	}
	return tmp
}

func removeElement(list []int, element int) []int {
	tmp := make([]int, 0)
	index, contem := containValueOnList(list, element)
	if contem {
		tmp = append(tmp, list[:index]...)
		tmp = append(tmp, list[index+1:]...)
		return tmp
	}
	return list
}

func removeLastElement(list []int) []int {
	if len(list) == 0 {
		return list
	}
	tmp := make([]int, 0)
	tmp = append(tmp, list[:(len(list)-1)]...)
	return tmp
}

func makeSet(n, size int, list []int) [][]int {
	nextSteps := calcNextSteps(n, list)
	for _, value := range nextSteps {
		list = append(list, value)
		if !testState(list, size) {
			list = removeElement(list, value)
		} else {
			makeSet(n, size, list)
			list = removeLastElement(list)
		}
	}
	return returnList
}

func callMinToMax(n, minSize, maxSize int) [][]int {
	var setList [][]int
	for size := minSize; size <= maxSize; size++ {
		setList = makeSet(n, size, []int{})
	}
	return setList
}
