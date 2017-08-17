package main

import "fmt"

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    var w func(t *tree.Tree)
    w = func(t *tree.Tree) {
        if t.Left != nil {
            w(t.Left)
        }
        ch <-t.Value
        if t.Right != nil {
            w(t.Right)
        }
    }
    w(t)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int)
    c2 := make(chan int)

    go Walk(t1, c1)
    go Walk(t2, c2)

    for {
        n1, ok1 := <-c1
        fmt.Println("n1", n1)
        n2, ok2 := <-c2
        fmt.Println("n2", n2)
        if n1 != n2{
            fmt.Println("[Info] n1 != n2")
            return false
        }
        if ok1 != ok2 {
            fmt.Println("[Info] ok1 != ok2")
            return false
        }
        if !ok1 {
            fmt.Println("[Info] !ok1")
            break
        }
    }
    return true
}

func main() {
    fmt.Println("***")
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
