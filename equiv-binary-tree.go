package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1_chan := make(chan int)
	t2_chan := make(chan int)

	go Walk(t1, t1_chan)
	go Walk(t2, t2_chan)

	for i := 0; i < 10; i++ {
		t1_e := <-t1_chan
		t2_e := <-t2_chan

		if t1_e != t2_e {
			return false
		}
	}

	return true
}

func main() {
	// test Walk()
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 1; i <= 10; i++ {
		ii := <-ch
		if ii != i {
			fmt.Println("Walk Error, doesn't match!")
		}
	}

	// test Same()
	fmt.Println("Same(tree.New(1), tree.New(1)): ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same(tree.New(1), tree.New(2)): ", Same(tree.New(1), tree.New(2)))
}
