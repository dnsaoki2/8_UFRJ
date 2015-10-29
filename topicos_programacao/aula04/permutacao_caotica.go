//entrada: numero inteiro n
//saida: permutaçoes caoticas de {0,...,n} (o numero nao pode estar na sua posicao)
//exemplo: n = 2 {0,1,2}
//saida: {1,2,0}
// {2,0,1}

package main

import "fmt"

var list = make([]int, 0)

func testState(list []int) bool {
	for index, value := range list {
		if value == index {
			return false
		}
	}
	return true
}

func contain(list []int, x int) (int, bool) {
	for index, value := range list {
		if value == x {
			return index, true
		}
	}
	return 0, false
}

func calcNextSteps(n int, list []int) []int {
	tmp := make([]int, 0)
	for i := 0; i < n; i++ {
		if _, ok := contain(list, i); !ok {
			tmp = append(tmp, i)
		}
	}
	return tmp
}

func removeElement(list []int, element int) []int {
	tmp := make([]int, 0)
	index, contem := contain(list, element)
	if contem {
		tmp = append(tmp, list[:index]...)
		tmp = append(tmp, list[index+1:]...)
		return tmp
	}
	return nil
}

func removeLastElement(list []int) []int {
	if len(list) == 0 {
		return list
	}
	tmp := make([]int, 0)
	tmp = append(tmp, list[:(len(list)-1)]...)
	return tmp
}

func backtrack(n int) {
	nextSteps := calcNextSteps(n, list)
	for _, value := range nextSteps {
		list = append(list, value)
		if !testState(list) {
			list = removeElement(list, value)
		} else {
			if len(list) == n && testState(list) {
				fmt.Println(list)
				list = removeLastElement(list)
			} else {
				backtrack(n)
			}
		}
	}
	list = removeLastElement(list)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	backtrack(n)
}
