// Package main
package main

import (
	"fmt"
	"strings"
	"unicode"
)

const punctuation string = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~."

func main() {
	var value string
	value = "why so serious?"

	fmt.Println(FirstLetterUp(value))
}

// FirstLetterUp Функция приниает на вход переменную value типа string
// и возвращает преобразованую строку, где первая буква заглавная, а на конце стоит точка, если знака пунктуации нет.
// FirstLetterUp("hello world") --> Hello world.
func FirstLetterUp(value string) string {
	chars := []rune(value)
	lenght := len(chars)
	if lenght > 0 {
		first := unicode.ToUpper(chars[0])
		last := string(chars[lenght-1])
		result := string(first) + string(chars[1:])
		if strings.Contains(punctuation, last) != true {
			result += "."
		}
		return result
	}
	return value
}
