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
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int64](10)
	if _, ok := s.Peek(); ok {
		t.Errorf("expected stack.Peek on empty stack to return empty stack error, got %v instead", ok)
	}
	for i := 1; i <= 10; i++ {
		ok := s.Push(int64(i))
		if !ok {
			t.Error("failed to push to stack")
		}
	}
	v, ok := s.Peek()
	if v != 10 || !ok {
		t.Errorf("expected stack.Peek to return value 10, got %v instead and ok = %v", v, ok)
	}
	if !s.IsFull() || s.Size() != 10 {
		t.Errorf("expected stack of capacity 10 to be full, have size %v", s.Size())
	}

	ok = s.Push(11)
	if ok {
		t.Errorf("expected false, got %v instead", ok)
	}
	for i := 10; i > 0; i-- {
		v, ok := s.Pop()
		if v != int64(i) || !ok {
			t.Errorf("expected %v got %v instead with ok=%v", i, v, ok)
		}
	}
	_, ok = s.Pop()
	if ok {
		t.Errorf("expected false got %v instead", ok)
	}
}
