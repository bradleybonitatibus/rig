# `generator`

This package contains a "generator" implementation allowing generation of values
that can be processed iteratively.

This implementation uses Go's `generics`, and the return type of the 
generation function is used as the value that is sent over the read-only
channel that yields values from it.


```go
package main

import (
	"fmt"
	"math/rand"

    "github.com/bradleybonitatibus/rig/generator"
)

func main() {
	g := generator.New(100, func() string {
		return fmt.Sprintf("hello_%v", rand.Int())
	})

	for v := range g.Iter() {
		fmt.Println(v)
		// Output: "hello_"
	}
}
```