package greeter

import (
	"fmt"
	"strings"
)

func Greet(name string, hour int) string {

	if hour < 0 || hour > 23 {
		return "Hour out of range"
	} else {
		var greeting string

		if hour >= 6 && hour < 12 {
			greeting = "Good morning"
		} else if hour >= 12 && hour <= 18 {
			greeting = "Hello"
		} else if hour >= 18 && hour < 22 {
			greeting = "Good evening"
		} else if hour >= 22 && hour < 24 || hour >= 0 && hour < 6 {
			greeting = "Good night"
		} else {

		}
		trimmedName := strings.Trim(name, " ")
		return fmt.Sprintf("%s, %s!", greeting, strings.Title(trimmedName))
	}
}
