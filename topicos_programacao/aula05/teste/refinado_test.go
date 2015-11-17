package main

import (
	"gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type testConjunto struct{}

var _ = check.Suite(&testConjunto{})

func (s *testConjunto) SetUpTest(c *check.C) {
}

func (t *testConjunto) TestState(c *check.C) {
	arrayList := [][]int{}
	arrayListResult, result := testState([]int{1}, arrayList, 1)
	c.Assert(result, check.Equals, false)
	c.Assert(arrayListResult, check.DeepEquals, [][]int{[]int{1}})
	arrayListResult, result = testState([]int{}, arrayList, 1)
	c.Assert(result, check.Equals, true)
	c.Assert(arrayListResult, check.DeepEquals, [][]int{})
	arrayListResult, result = testState([]int{1, 2}, arrayList, 1)
	c.Assert(result, check.Equals, false)
	c.Assert(arrayListResult, check.DeepEquals, [][]int{})
}

func (t *testConjunto) TestContain(c *check.C) {
	index, result := containValueOnList([]int{1, 2, 3, 4}, 3)
	c.Assert(result, check.Equals, true)
	c.Assert(index, check.Equals, 2)
	index, result = containValueOnList([]int{1, 2, 3, 4}, 5)
	c.Assert(result, check.Equals, false)
	c.Assert(index, check.Equals, 0)
}

func (t *testConjunto) TestContainArrayList(c *check.C) {
	arrayList := [][]int{[]int{1}, []int{2, 1}, []int{2, 3}}
	result := containListOnArrayList(arrayList, []int{1})
	c.Assert(result, check.Equals, true)
	result = containListOnArrayList(arrayList, []int{2, 1})
	c.Assert(result, check.Equals, true)
	result = containListOnArrayList(arrayList, []int{1, 2})
	c.Assert(result, check.Equals, true)
	result = containListOnArrayList(arrayList, []int{1, 4})
	c.Assert(result, check.Equals, false)
}

func (t *testConjunto) TestCalcNextSteps(c *check.C) {
	result := calcNextSteps(2, []int{})
	c.Assert(len(result), check.Equals, 2)
	c.Assert(result, check.DeepEquals, []int{1, 2})
	result = calcNextSteps(2, []int{1})
	c.Assert(len(result), check.Equals, 1)
	c.Assert(result, check.DeepEquals, []int{2})
}

func (t *testConjunto) TestRemoveElement(c *check.C) {
	result := removeElement([]int{1, 2, 3}, 2)
	c.Assert(result, check.DeepEquals, []int{1, 3})
	result = removeElement([]int{1, 3}, 2)
	c.Assert(result, check.DeepEquals, []int{1, 3})
}

func (t *testConjunto) TestRemoveLastElement(c *check.C) {
	result := removeLastElement([]int{1, 2, 3})
	c.Assert(result, check.DeepEquals, []int{1, 2})
	result = removeLastElement([]int{})
	c.Assert(result, check.DeepEquals, []int{})
}

func (t *testConjunto) TestCallMinToMax(c *check.C) {
	result := callMinToMax(4, 2, 3)
	expected := [][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 3}, []int{2, 4}, []int{3, 4},
		[]int{1, 2, 3}, []int{1, 2, 4}, []int{1, 3, 4}, []int{2, 3, 4}}
	c.Assert(result, check.DeepEquals, expected)
}
