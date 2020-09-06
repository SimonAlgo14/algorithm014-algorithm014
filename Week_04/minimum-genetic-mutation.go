package main

import (
	"fmt"
)


func minMutation(start string, end string, bank []string) int {
	queue   := []string{ start}
	genmap  := make(map[string]bool)
	visited := make(map[string]bool)

	for _, gen := range bank {
		genmap[gen] = true
	}

	count := 0
	for len(queue) > 0 {
		var kids []string
		for _, one := range queue {
			if one == end {
				return count
			}

			curr := []rune(one)
			for i := 0; i < len(curr); i++ {
				old := curr[i]

				for _, g := range "ACGT" {
					curr[i] = g
					newGen := string(curr)

					_, okV := visited[newGen]
					_, okG := genmap[newGen]
					if ! okV && okG {
						kids = append(kids, newGen)
						visited[newGen] = true

						fmt.Printf("count: %d change i: %d %c %s ->  %s \n", count, i, g, one, newGen)
					}
				}

				curr[i] = old
			}
		}

		queue = kids
		count++
	}

	return -1
}



func main() {
	fmt.Println(minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}))
}

/*


Example 1:

start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]

return: 1



Example 2:

start: "AACCGGTT"
end:   "AAACGGTA"
bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]

AACCGGTT
  *
AAACGGTT
       *
AAACGGTA

return: 2



Example 3:

start: "AAAAACCC"
end:   "AACCCCCC"
bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]

return: 3


*/
