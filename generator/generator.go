/*
Copyright 2022 Bradley Bonitatibus

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generator

// GeneratorFn is a generic function to returns a value of type T.
type GeneratorFn[T any] func() T

// Generator is a concept that will allow users to define a function to return
// values of type T for a specified size amount.
type Generator[T any] struct {
	fn   GeneratorFn[T]
	size int
}

// New is the factory method to create a new Generator. The number of values that
// can be generated from this function is controlled by the size parameter.
// fn is the function that returns a given value of type T.
func New[T any](size int, fn GeneratorFn[T]) *Generator[T] {
	return &Generator[T]{
		fn:   fn,
		size: size,
	}
}

// Iter returns a read-only channel that will have values sent through it
// from a separate go routine.
func (g *Generator[T]) Iter() <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)
		for i := 0; i < g.size; i++ {
			c <- g.fn()
		}
	}()

	return c
}
