package grammer

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	s := "babad"
	palindrome := longestPalindrome(s)
	fmt.Println(palindrome)
}
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	var ans string
	for l := 0; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}

			if dp[i][j] && (l+1 > len(ans)) {
				ans = s[i : j+1]
			}

		}
	}
	return ans
}
