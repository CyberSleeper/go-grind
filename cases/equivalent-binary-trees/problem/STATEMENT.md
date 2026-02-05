## Problem: Equivalent Binary Trees

### 📌 Concept

In Go, concurrency primitives (Channels and Goroutines) can be used to solve algorithmic problems in elegant ways. This problem asks you to determine if two Binary Search Trees hold the same sequence of values, regardless of their structure.

This is a classic exercise from "A Tour of Go" that demonstrates:

1.  **Tree Traversal** (Recursion).
2.  **Channels** as pipes/streams of data.
3.  **Synchronization** (ensuring one tree doesn't race ahead of the other improperly).

### 📝 Task

1.  Implement `Walk(t *tree.Tree, ch chan int)`.
    - It should traverse the tree `t` in-order (Left, Root, Right).
    - Send each value to channel `ch`.
    - **Crucial:** Close the channel when traversal is finished.
2.  Implement `Same(t1, t2 *tree.Tree) bool`.
    - It should determine whether the trees `t1` and `t2` contain the same values in the same order.
    - It must use `Walk` to traverse both trees concurrently.
    - It should return `true` if sequences match, `false` otherwise.

### 🚫 Starter Code

```go
package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    // TODO: Implement In-Order Traversal
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    // TODO: Run Walk on t1 and t2 concurrently
    // Compare the values received from both channels
    return false
}

func main() {
    // Should print true
    fmt.Println(Same(tree.New(1), tree.New(1)))
    // Should print false
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
```

### 💡 Hint

You will likely need a helper function for the recursion inside Walk so that you can close the channel exactly once in the main Walk function.
