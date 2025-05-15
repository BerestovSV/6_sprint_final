package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}
func isRussianLetter(r rune) bool {
	upper := unicode.ToUpper(r)
	return upper >= 'А' && upper <= 'Я' || upper == 'Ё'
}

func AutoConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", errors.New("передана пустая строка")
	}

	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	for _, r := range trimmed {
		if unicode.IsLetter(r) && !isRussianLetter(r) {
			return "", errors.New("обнаружены символы не из русского алфавита")
		}
	}

	return morse.ToMorse(trimmed), nil
}
