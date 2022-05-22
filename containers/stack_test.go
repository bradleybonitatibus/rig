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
	"errors"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int64](10)
	if _, err := s.Peek(); err == nil {
		t.Errorf("expected stack.Peek on empty stack to return empty stack error, got %v instead", err)
	}
	for i := 1; i <= 10; i++ {
		err := s.Push(int64(i))
		if err != nil {
			t.Error(err)
		}
	}
	v, err := s.Peek()
	if v != 10 || err != nil {
		t.Errorf("expected stack.Peek to return value 10, got %v instead and err = %v", v, err)
	}
	if !s.IsFull() || s.Size() != 10 {
		t.Errorf("expected stack of capacity 10 to be full, have size %v", s.Size())
	}

	err = s.Push(11)
	if !errors.Is(err, ErrStackFull) {
		t.Errorf("expected error capacity limit reached, got %v instead", err)
	}
	for i := 10; i > 0; i-- {
		v, err := s.Pop()
		if v != int64(i) || err != nil {
			t.Errorf("expected %v got %v instead with err=%v", i, v, err)
		}
	}
	_, err = s.Pop()
	if err != ErrEmptyStack {
		t.Errorf("expected empty stack, got %v instead", err)
	}
}
