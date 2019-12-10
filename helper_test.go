package main

import (
	"fmt"
	"testing"
	"time"
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

func TestMorning(t *testing.T) {
	sample := time.Date(0, 0, 0, 10, 15, 30, 918273645, time.UTC)
	ans := isMorning(sample)
	part := partOfTheDay(sample)
	fmt.Println(part)
	if ans != true || part != "morning" {
		t.Errorf("Time %s is not morning", sample.Local())
	}
}

func TestAfternoon(t *testing.T) {
	sample := time.Date(0, 0, 0, 19, 15, 30, 918273645, time.UTC)
	ans := isAfternoon(sample)
	part := partOfTheDay(sample)
	fmt.Println(part)
	if ans != true || part != "afternoon" {
		t.Errorf("Time %s is not afternoon", sample.Local())
	}
}

func TestEvening(t *testing.T) {
	sample := time.Date(0, 0, 0, 23, 15, 30, 918273645, time.UTC)
	ans := isEvening(sample)
	part := partOfTheDay(sample)
	if ans != true || part != "evening" {
		t.Errorf("Time %s is not evening", sample.Local())
	}
}
