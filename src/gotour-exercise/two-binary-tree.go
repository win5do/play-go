package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func _walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	_walk(t.Left, ch)
	_walk(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	res := true

	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2

		res = x == y && ok1 == ok2 != false

		if x != y || ok1 == false || ok2 == false {
			return res
		}
	}

}

func main() {
	t1 := tree.New(2)
	t2 := tree.New(3)
	r := Same(t1, t2)
	fmt.Println(r)
}
