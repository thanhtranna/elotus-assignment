package maximumlengthofrepeatedsubarray

func findLength(A []int, B []int) int {
	lengthArrOne, lengthArrTwo := len(A), len(B)

	// dp[i][j] == k means A[i-k:i] == B[j-k:j] ,
	// but A[i-k-1] != B[j-k-1]
	dp := make([][]int, lengthArrOne+1)
	for i := range dp {
		dp[i] = make([]int, lengthArrTwo+1)
	}

	result := 0
	for i := 1; i <= lengthArrOne; i++ {
		for j := 1; j <= lengthArrTwo; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				result = max(result, dp[i][j])
			}
		}
	}

	return result
}

func max(numOne, numTwo int) int {
	if numOne > numTwo {
		return numOne
	}

	return numTwo
}
