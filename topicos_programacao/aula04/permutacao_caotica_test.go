package main

import "testing"

func TestCalcNextSteps(t *testing.T) {
	result := calcNextSteps(2, []int{1})
	expected := []int{0, 2}
	if result[0] != expected[0] && result[1] != expected[1] {
		t.Error("Esperado {0,2}, obteve ", result)
	}
}

func TestContain(t *testing.T) {
	_, result := contain([]int{1, 2, 3}, 1)
	_, result2 := contain([]int{1, 2, 3}, 4)
	if result != true || result2 != false {
		t.Error("Esperado true, false, obteve:  ", result, result2)
	}
}

func TestState(t *testing.T) {
	result := testState([]int{0, 1, 2})
	result2 := testState([]int{1, 2, 0})
	if result != false || result2 != true {
		t.Error("Esperado false,true, obteve:  ", result, result2)
	}
}

func TestRemoveElement(t *testing.T) {
	result := removeElement([]int{1, 2, 3}, 2)
	result2 := removeElement([]int{1, 2, 3}, 5)
	if result[0] != 1 || result[1] != 3 || result2 != nil {
		t.Error("Esperado 1,3,nil obteve:  ", result, result2)
	}
}

func TestRemoveLastElement(t *testing.T) {
	result := removeLastElement([]int{1, 2, 3})
	if result[0] != 1 || result[1] != 2 {
		t.Error("Esperado 1,3,nil obteve:  ", result)
	}
}
