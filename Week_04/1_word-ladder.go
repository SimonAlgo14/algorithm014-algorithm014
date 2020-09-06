package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) <= 0 {
		return 0
	}

	queue := []string{beginWord }
	nextQueue := make([]string, 0)
	wordHash := make(map[string]bool)
	count   := 1

	for _, w := range wordList {
		wordHash[w] = true
	}

	for len(queue) > 0 {
		nextQueue = nextQueue[:0]

		for _, one := range queue {
			if one == endWord {
				return count
			}

			curr := []rune(one)
			for i := 0; i < len(curr); i++ {
				old := curr[i]

				for _, letter := range letters() {
					curr[i] = letter
					nextCurr := string(curr)

					if _, ok := wordHash[nextCurr]; ok {
						nextQueue = append(nextQueue, nextCurr)
						delete(wordHash, nextCurr)

						//fmt.Printf("%d add %s -> %s \n", count, one, nextCurr)
					}
				}

				curr[i] = old
			}
		}

		queue = append(queue[:0], nextQueue...)
		count++
	}

	return 0
}

func letters() []rune{
	r := make([]rune, 64)
	for char := 'a'; char <= 'z'; char ++ {
		r = append(r, char)
	}

	return r
}

func main() {
	words := []string{"ymann", "yycrj", "oecij", "ymcnj", "yzcrj", "yycij", "xecij", "yecij", "ymanj", "yzcnj", "ymain"}
	fmt.Println(ladderLength("ymain", "oecij", words))
}


