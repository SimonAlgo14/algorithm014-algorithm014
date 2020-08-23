package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func buildBST(a []int) *Tree {
	if len(a) <= 0 {
		return nil
	}

	sort.Ints(a)
	return build(a, 0, len(a)-1)
}

func build(a []int, low int, high int) *Tree {
	if low > high {
		return nil
	}

	mid := low + (high-low)/2
	node := &Tree{Value: a[mid]}
	node.Left = build(a, low, mid-1)
	node.Right = build(a, mid+1, high)

	return node
}

func visit(tree *Tree) {
	if tree == nil {
		return
	}

	visit(tree.Left)
	fmt.Printf("%d ", tree.Value)
	visit(tree.Right)
}

func visit2(tree *Tree) {
	if tree == nil {
		return
	}

	visit2(tree.Right)
	fmt.Printf("%d ", tree.Value)
	visit(tree.Left)
}


func visit3(tree *Tree, level int, flag string, root string) {
	if tree == nil {
		return
	}

	fmt.Printf("%s%d (%s -> %s)\n", strings.Repeat("|  ", level), tree.Value, flag, root)
	visit3(tree.Left, level + 1, "l", strconv.Itoa(tree.Value))
	visit3(tree.Right, level + 1, "r", strconv.Itoa(tree.Value))
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8,9}
	t := buildBST(a)
	visit(t)

	fmt.Println()
	visit2(t)


	fmt.Println()
	visit3(t, 0, "r", "nil")
}
