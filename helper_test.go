package main

import (
	"testing"
)

func TestIsNotPalindrome(t *testing.T) {
	// ans := IntMin(2, -2)
	noPalindrome := "nothing"
	ans := isPalindrome(noPalindrome)
	if ans == true {
		t.Errorf("%s is not a palindrome", noPalindrome)
	}
}

func TestIsPalindromeEvenTerm(t *testing.T) {
	palindrome := "aabbaa"
	ans := isPalindrome(palindrome)
	if ans == false {
		t.Errorf("%s is a palindrome", palindrome)
	}
}

func TestIsPalindromeOddTerm(t *testing.T) {
	palindrome := "aerea"
	ans := isPalindrome(palindrome)
	if ans == false {
		t.Errorf("%s is a palindrome", palindrome)
	}
}

func TestIsEmptyPalindrome(t *testing.T) {
	palindrome := ""
	ans := isPalindrome(palindrome)
	if ans == false {
		t.Errorf("%s is a palindrome", palindrome)
	}
}

func TestReverse(t *testing.T) {
	sample := "sample"
	expected := "elpmas"
	ans := reverse(sample)
	if ans != expected {
		t.Errorf("%s not the reverse of %s (%s)", ans, sample, expected)
	}
}

func TestReversePalindrome(t *testing.T) {
	sample := "aerea"
	ans := reverse(sample)
	if sample != ans {
		t.Errorf("%s is not the reverse of %s", ans, sample)
	}
}

func TestAnswerMorning(t *testing.T) {

}

func TestAnswerAfternoon(t *testing.T) {

}

func TestAnswerEvening(t *testing.T) {

}
