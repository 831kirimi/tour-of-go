package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
    walk(t, ch)
    close(ch)
}

func walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }

    walk(t.Left, ch)
    ch <- t.Value
    walk(t.Right, ch)
}



// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	v1 := <- ch1
	v2 := <- ch2
	if v1 == v2{
	return true	
	}else{
		return false
	}
}


func main() {
	ch := make(chan int)
    go Walk(tree.New(1), ch)
	for i := range ch {
        fmt.Println(i)
    }

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
    
}
