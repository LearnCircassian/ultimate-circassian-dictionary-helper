package utils

import (
	"strings"
	"unicode"
)

func IsLatin(word string) bool {
	return !IsCyrillic(word)
}

func IsCyrillic(word string) bool {
	word = strings.Trim(word, " ")

	if len(word) < 1 {
		return false
	}

	if word[0] == '1' {
		return true
	}

	text := []rune(word)
	if len(text) < 1 {
		return false
	}

	// check first letter
	if unicode.Is(unicode.Cyrillic, text[0]) {
		return true
	}

	return false
}
