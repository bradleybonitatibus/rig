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

package containers

import (
	"sync"
)

// Stack data structure implementation that is thread safe.
type Stack[T any] struct {
	mu       sync.RWMutex
	capacity int
	values   []T
	size     int
}

// NewStack creates an empty stack with a predefined capacity.
func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		mu:       sync.RWMutex{},
		values:   make([]T, 0, capacity),
		capacity: capacity,
		size:     0,
	}
}

// Pop returns the top value in the stack, or an error if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsEmpty() {
		var empty T
		return empty, false
	}
	last := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	s.size--
	return last, true
}

// Push an item to the top of the stack. If the stack has reached it's capacity,
// it will return ErrStackFull.
func (s *Stack[T]) Push(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.IsFull() {
		return false
	}
	s.size++
	s.values = append(s.values, item)
	return true
}

// Peek returns the value at the top of the stack. If the stack is empty,
// it will return the "default" value of type T, and false.
func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.IsEmpty() {
		var empty T
		return empty, false
	}
	return s.values[len(s.values)-1], true
}

// IsEmpty returns true when the stack does not contain any values.
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

// IsFull returns true when the stack has reached full capacity.
func (s *Stack[T]) IsFull() bool {
	return s.size == s.capacity
}

// Size returns the current size of the stack.
func (s *Stack[T]) Size() int {
	return s.size
}
