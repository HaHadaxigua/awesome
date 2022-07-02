package prefix_sum

import (
	"awesome/util"
	"fmt"
	"math"
	"testing"
)

/**
Given a binary array nums,

return the maximum length of a contiguous subarray with an equal number of 0 and 1.
*/
func findMaxLength(nums []int) int {
	// check if current array with an equal number of 0 and 1
	var (
		count  int
		maxLen int = math.MinInt32
	)

	mem := make(map[int]int)
	mem[0] = -1
	for i, n := range nums {
		if n == 1 {
			count++
		} else {
			count--
		}

		if v, ok := mem[count]; ok {
			maxLen = util.Max(maxLen, i-v)
		} else {
			mem[count] = i
		}
	}
	return maxLen
}

func TestFindMaxLength(t *testing.T) {
	nums := []int{0, 1, 1, 0}
	length := findMaxLength(nums)
	fmt.Println(length)
}
