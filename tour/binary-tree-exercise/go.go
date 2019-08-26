package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {

	// descend as far as possible to the left
	if t.Left != nil {
		walkTree(t.Left, ch)
	}

	// at this point t.Left is nil
	// we're either a leaf node or a node with only a right node
	// thus we're the next node to be printed, send it to the channel
	ch <- t.Value

	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int)
	go Walk(tree.New(2), c)

	for i := range c {
		println(i)
	}

	println(Same(tree.New(1), tree.New(1)))
	println(Same(tree.New(1), tree.New(2)))

}

/*
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
} */
