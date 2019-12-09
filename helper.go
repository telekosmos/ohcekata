package main

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	c, l := len(runes)/2, len(runes)

	for i := 0; i < c; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
