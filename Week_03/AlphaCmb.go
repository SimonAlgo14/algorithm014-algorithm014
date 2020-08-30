package main

import "fmt"

func letterCombinations(digits string) []string {
	dict := map[rune] string {
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var rs []string
	letterCmb(dict, digits, 0, len(digits)-1, []rune{}, &rs)
	return rs
}


func letterCmb(dict map[rune] string,  digits string, index int, end int, cache []rune, rs *[]string)  {
	if index > end {
		*rs = append(*rs, string(cache))
		return
	}

	letters := dict[ []rune(digits)[index] ]

	if len(letters) == 0 {
		letterCmb(dict, digits, index + 1, end, cache, rs)
	} else {
		for _, letter := range letters {
			letterCmb(dict, digits, index + 1, end, append(cache, letter), rs)
		}
	}
}

func main() {
	fmt.Println(letterCombinations( "23"))
	fmt.Println(letterCombinations( "123"))
	fmt.Println(letterCombinations( "231"))
	fmt.Println(letterCombinations( "2319"))
}