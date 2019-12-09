package main

import (
	"testing"
)

func TestIsNotPalindrome(t *testing.T) {
	// ans := IntMin(2, -2)
	noPalindrome := "nothing"
	ans := isPalindrome(noPalindrome)
	if ans == true {
		t.Errorf("%t is not a palindrome", ans)
	}
}

func TestIsPalindromeEvenTerm(t *testing.T) {

}

func TestIsPalindromeOddTerm(t *testing.T) {

}

func TestIsEmptyPalindrome(t *testing.T) {

}

func TestReverse(t *testing.T) {

}

func TestReversePalindrome(t *testing.T) {

}

func TestAnswerMorning(t *testing.T) {

}

func TestAnswerAfternoon(t *testing.T) {

}

func TestAnswerEvening(t *testing.T) {

}
