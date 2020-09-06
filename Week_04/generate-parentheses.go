package main

import "fmt"

func generateParenthesis(n int) []string {
	var r []string
	dfs(&r, n, 0, 0, "")
	return r
}

func dfs(r *[]string, n int, left int, right int, path string) {
	if n == left && n == right {
		*r = append(*r, path)
		return
	}

	if left < n {
		dfs(r, n, left + 1, right, path + "(")
	}
	if right < left {
		dfs(r, n, left, right + 1, path + ")")
	}
}

func main() {
	fmt.Println( generateParenthesis(2) )
	fmt.Println( generateParenthesis(3) )
	fmt.Println( generateParenthesis(4) )
}
