package main

import "fmt"

func main() {
	s := "abc"
	e := "def"
	t := fmt.Sprintf("%c%s", []rune(e)[0], string([]rune(s)[1:]))
	fmt.Println(t)
}
