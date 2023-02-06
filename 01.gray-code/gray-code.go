package graycode

func grayCode(n int) []int {
	var length uint = 1 << uint(n)
	out := make([]int, length)

	for i := uint(0); i < length; i++ {
		out[i] = int((i >> 1) ^ i)
	}

	return out
}
