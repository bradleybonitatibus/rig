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

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	type user struct {
		id        string
		purchases []string
	}

	stringGenFn := func() string {
		return fmt.Sprintf("hello_%v", rand.Int())
	}

	structGenFn := func() user {
		return user{
			id:        fmt.Sprintf("user_id_%v", rand.Int()),
			purchases: []string{"order_id", "other_order_id"},
		}
	}

	sGen := New(10, stringGenFn)
	if sGen.size != 10 || sGen.fn == nil {
		t.Error("failed to set size for string generator")
	}

	userGen := New(20, structGenFn)
	if userGen.size != 20 || userGen.fn == nil {
		t.Error("failed to set size for user generator")
	}

	sCount := 0
	for s := range sGen.Iter() {
		if s == "" {
			t.Error("string generator yielded empty string")
		}
		sCount++
	}

	if sCount != 10 {
		t.Errorf("expected 10, got %v instead", sCount)
	}
	uCount := 0
	for u := range userGen.Iter() {
		if u.id == "" || len(u.purchases) == 0 {
			t.Errorf("userGen yielded empty value")
		}
		uCount++
	}
	if uCount != 20 {
		t.Errorf("expected 20 users to be generated, got %v instead", uCount)
	}
}
