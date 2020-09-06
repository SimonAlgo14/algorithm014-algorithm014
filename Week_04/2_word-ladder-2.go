package main

type item struct {
	word string
	path []string
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	_, find := Find(wordList, endWord)
	if !find || len(beginWord) == 0 || len(endWord) == 0 {
		return make([][]string, 0)
	}

	n := len(beginWord)
	combo := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < n; i++ {
			iter_w := w[:i] + "*" + w[i+1:]
			combo[iter_w] = append(combo[iter_w], w)
		}
	}

	queue := make([]item, 0)
	queue = append(queue, item{beginWord, []string{beginWord}})
	visited := make(map[string]bool)
	visited[beginWord] = true

	res := make([][]string, 0)
	for len(queue) > 0 {
		k := len(queue)
		next_layer_visited := make(map[string]bool)
		for i := 0; i < k; i++ {
			pop := queue[0]
			curr_word := pop.word
			path := pop.path

			if curr_word == endWord {
				res = append(res, path)
			}

			for j := 0; j < n; j++ {
				iter_w := curr_word[:j] + "*" + curr_word[j+1:]
				_, ok := combo[iter_w]
				if ok {
					next_words := combo[iter_w]
					for _, next_word := range next_words {
						_, ok_visited := visited[next_word]
						if !ok_visited {
							var new_path = make([]string, len(path))
							copy(new_path, path)
							new_path = append(new_path, next_word)
							queue = append(queue, item{next_word, new_path})
							next_layer_visited[next_word] = true
						}
					}
				}
			}
			queue = queue[1:]
		}
		if len(res) > 0 {
			return res
		}
		for key := range next_layer_visited {
			visited[key] = true
		}
	}

	return res
}


func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
