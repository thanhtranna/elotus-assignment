package sumofdistancesintree

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
	N     int
	edges [][]int
}

// answer is the answer
// one represents the first answer
type answer struct {
	one []int
}

func Test_SumOfDistancesInTree(t *testing.T) {
	testCases := []question{
		{
			param{4, [][]int{{1, 2}, {3, 2}, {3, 0}}},
			answer{[]int{6, 6, 4, 4}},
		},
		{
			param{6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}},
			answer{[]int{8, 12, 6, 10, 10, 10}},
		},
	}

	fmt.Printf("------------------------ Run tests ------------------------\n")

	for _, test := range testCases {
		result := sumOfDistancesInTree(test.param.N, test.param.edges)
		ok := reflect.DeepEqual(result, test.answer.one)

		if !ok {
			fmt.Printf("【input】:%v【output】:%v\n", test.param, result)
		}
	}
	fmt.Printf("\n\n\n")
}
