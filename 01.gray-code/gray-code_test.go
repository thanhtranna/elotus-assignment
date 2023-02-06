package graycode

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
	one int
}

// answer is the answer
// one represents the first answer
type answer struct {
	one []int
}

func Test_GrayCode(t *testing.T) {
	testCases := []question{
		{
			param{2},
			answer{[]int{0, 1, 3, 2}},
		},

		{
			param{0},
			answer{[]int{0}},
		},

		{
			param{3},
			answer{[]int{0, 1, 3, 2, 6, 7, 5, 4}},
		},
	}

	fmt.Printf("------------------------ Run tests ------------------------\n")

	for _, test := range testCases {
		result := grayCode(test.param.one)
		ok := reflect.DeepEqual(result, test.answer.one)

		if !ok {
			fmt.Printf("【input】:%v【output】:%v\n", test.param, result)
		}
	}
	fmt.Printf("\n\n\n")
}
