package main

import (
	"time"
)

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

func isMorning(t time.Time) (between bool) {
	init := 6
	end := 12
	hour, minute, second := time.Time.Hour(t), time.Time.Minute(t), time.Time.Second(t)
	between = hour >= init || (hour <= end && minute == 0 && second == 0)

	return
}

func isAfternoon(t time.Time) (between bool) {
	init := 12
	end := 20
	hour, minute, second := time.Time.Hour(t), time.Time.Minute(t), time.Time.Second(t)
	between = hour >= init || (hour <= end && minute == 0 && second == 0)

	return
}

func isEvening(t time.Time) (between bool) {
	init := 20
	end := 6
	hour, minute, second := time.Time.Hour(t), time.Time.Minute(t), time.Time.Second(t)
	between = hour >= init || (hour <= end && minute == 0 && second == 0)

	return
}

func partOfTheDay(t time.Time) (part string) {
	eveningMorning := 6
	morningAfternoon := 12
	afternoonEvening := 20
	hour, minute, second := time.Time.Hour(t), time.Time.Minute(t), time.Time.Second(t)
	morning := hour >= eveningMorning && (hour <= morningAfternoon || minute == 0 || second == 0)
	afternoon := hour >= morningAfternoon && (hour <= afternoonEvening || minute == 0 || second == 0)

	if morning == true {
		part = "morning"
	} else if afternoon == true {
		part = "afternoon"
	} else {
		part = "evening"
	}

	return
}

func greet() (ans string) {
	m := make(map[string]string)
	m["morning"] = "Buenos días"
	m["afternoon"] = "Buenas tardes"
	m["evening"] = "Buenas noches"

	ans = m[partOfTheDay(time.Now())]

	return
}
