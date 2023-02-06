package maximumlengthofrepeatedsubarray

import (
	"fmt"
	"reflect"
	"testing"
)

type question struct {
	param
	answer
}

// param is the parameter
// one represents the first parameter
type param struct {
	A []int
	B []int
}

// answer is the answer
// one represents the first answer
type answer struct {
	one int
}

func Test_findLength(t *testing.T) {
	testCases := []question{
		{
			param{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
			answer{5},
		},
		{
			param{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}},
			answer{3},
		},
		{
			param{[]int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0}},
			answer{4},
		},
		{
			param{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}},
			answer{59},
		},
	}

	fmt.Printf("------------------------ Run tests ------------------------\n")

	for _, test := range testCases {
		result := findLength(test.param.A, test.param.B)
		ok := reflect.DeepEqual(result, test.answer.one)

		if !ok {
			fmt.Printf("【input】:%v【output】:%v\n", test.param, result)
		}
	}
	fmt.Printf("\n\n\n")
}
