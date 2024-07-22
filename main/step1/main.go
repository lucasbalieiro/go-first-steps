package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(len("four"))

	fmt.Println(len("fours"))

	fmt.Println(utf8.RuneCountInString("©"))

	fmt.Println(utf8.RuneCountInString("©©"))

}
