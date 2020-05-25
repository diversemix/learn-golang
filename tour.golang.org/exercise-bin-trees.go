package main

import (
	"fmt"
	"reflect"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func TreeToArray(t1 *tree.Tree) []int {
	t1ch := make(chan int)
	go Walk(t1, t1ch)
	t1array := make([]int, 10)
	for i := 0; i < 10; i++ {
		t1array[i] = <-t1ch
	}
	sort.Ints(t1array)
	return t1array
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	t1array := TreeToArray(t1)
	t2array := TreeToArray(t2)
	fmt.Println(t1array)
	fmt.Println(t2array)
	return reflect.DeepEqual(t1array, t2array)
}

func main() {
	var same = Same(tree.New(1), tree.New(1))
	fmt.Println("Expecting true:", same)

	var diff = Same(tree.New(1), tree.New(2))
	fmt.Println("Expecting false:", diff)

}
