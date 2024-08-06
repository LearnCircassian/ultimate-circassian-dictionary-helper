package utils

import (
	"regexp"
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

func ConvertIToCirStick(str string) string {
	str = strings.ReplaceAll(str, "ӏ", "1")
	str = strings.ReplaceAll(str, "ӏ", "1")
	// small i
	str = strings.ReplaceAll(str, "iа", "1а")
	str = strings.ReplaceAll(str, "iо", "1о")
	str = strings.ReplaceAll(str, "iе", "1е")
	str = strings.ReplaceAll(str, "iя", "1я")
	str = strings.ReplaceAll(str, "iы", "1ы")
	str = strings.ReplaceAll(str, "iэ", "1э")
	str = strings.ReplaceAll(str, "iи", "1и")
	str = strings.ReplaceAll(str, "iу", "1у")
	str = strings.ReplaceAll(str, "кi", "к1")
	str = strings.ReplaceAll(str, "шi", "ш1")
	str = strings.ReplaceAll(str, "пi", "п1")
	str = strings.ReplaceAll(str, "цi", "ц1")
	str = strings.ReplaceAll(str, "сi", "с1")
	str = strings.ReplaceAll(str, "чi", "ч1")
	str = strings.ReplaceAll(str, "тi", "т1")
	str = strings.ReplaceAll(str, "лi", "л1")
	str = strings.ReplaceAll(str, "фi", "ф1")
	str = strings.ReplaceAll(str, "щi", "щ1")

	// upper I
	str = strings.ReplaceAll(str, "Iа", "1а")
	str = strings.ReplaceAll(str, "Iо", "1о")
	str = strings.ReplaceAll(str, "Iе", "1е")
	str = strings.ReplaceAll(str, "Iя", "1я")
	str = strings.ReplaceAll(str, "Iы", "1ы")
	str = strings.ReplaceAll(str, "Iэ", "1э")
	str = strings.ReplaceAll(str, "Iи", "1и")
	str = strings.ReplaceAll(str, "Iу", "1у")
	str = strings.ReplaceAll(str, "кI", "к1")
	str = strings.ReplaceAll(str, "шI", "ш1")
	str = strings.ReplaceAll(str, "пI", "п1")
	str = strings.ReplaceAll(str, "цI", "ц1")
	str = strings.ReplaceAll(str, "сI", "с1")
	str = strings.ReplaceAll(str, "чI", "ч1")
	str = strings.ReplaceAll(str, "тI", "т1")
	str = strings.ReplaceAll(str, "лI", "л1")
	str = strings.ReplaceAll(str, "фI", "ф1")
	str = strings.ReplaceAll(str, "щI", "щ1")

	str = strings.ReplaceAll(str, "Ӏа", "1а")
	str = strings.ReplaceAll(str, "Ӏо", "1о")
	str = strings.ReplaceAll(str, "Ӏе", "1е")
	str = strings.ReplaceAll(str, "Ӏя", "1я")
	str = strings.ReplaceAll(str, "Ӏы", "1ы")
	str = strings.ReplaceAll(str, "Ӏэ", "1э")
	str = strings.ReplaceAll(str, "Ӏи", "1и")
	str = strings.ReplaceAll(str, "Ӏу", "1у")
	str = strings.ReplaceAll(str, "кӀ", "к1")
	str = strings.ReplaceAll(str, "шӀ", "ш1")
	str = strings.ReplaceAll(str, "пӀ", "п1")
	str = strings.ReplaceAll(str, "цӀ", "ц1")
	str = strings.ReplaceAll(str, "сӀ", "с1")
	str = strings.ReplaceAll(str, "чӀ", "ч1")
	str = strings.ReplaceAll(str, "тӀ", "т1")
	str = strings.ReplaceAll(str, "лӀ", "л1")
	str = strings.ReplaceAll(str, "фӀ", "ф1")
	str = strings.ReplaceAll(str, "щӀ", "щ1")
	return str
}

// Function to check if a string is fully capitalized
func IsFullyCapitalized(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func StartsWithNumber(s string) bool {
	if len(s) < 1 {
		return false
	}
	return unicode.IsNumber(rune(s[0]))
}

func StartsWithSpecialCharacter(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) < 1 {
		return false
	}
	r := rune(s[0])
	return !unicode.IsLetter(r) || r == '_' || r == '—' || r == '…' || r == '♦' || r == '■' || r == '•'
}

// Function to trim leading and trailing spaces and slashes
func TrimSlashes(s string) string {
	return strings.Trim(s, " /")
}

// Function to remove suffixes consisting of dots, commas, and numbers
func RemoveSuffixes(s string) string {
	return regexp.MustCompile(`[.,\d]+$`).ReplaceAllString(s, "")
}
