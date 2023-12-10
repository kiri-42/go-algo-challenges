package main

import (
	"errors"
	"unicode"
)

func IsPalindrome(s string) (isPal bool, err error) {
	if s == "" {
		return false, nil
	}

	if !isAlphanumeric(s) {
		return false, errors.New("英数字以外が含まれています")
	}

	s = reverseString(s)

	return s == reverseString(s), nil
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if (!unicode.IsLetter(r) || !unicode.In(r, unicode.Latin)) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
