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

package utils

import "testing"

func TestMustLoadEnv_PanicsWithoutEnvVar(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected to panic")
		}
	}()
	MustGetEnv("ABCWEKOEKUB1939KASINIMKAS")
}

func TestMustLoadEnv_ReturnsValidEnvVar(t *testing.T) {
	v := MustGetEnv("POSTGRES_HOST")
	if v == "" {
		t.Errorf("Expected non-empty string, got %v", v)
	}
}

func TestEnvVarWithFallback_ReturnsFallbackValueWhenEnvVarMissing(t *testing.T) {
	v := GetEnvWithFallback("OK1JIJIMKMK21ML", "hello")
	if v != "hello" {
		t.Errorf("Expected fallback value, got %v instead", v)
	}
}

func TestEnvVarWithFallback_ReturnsEnvVar(t *testing.T) {
	v := GetEnvWithFallback("POSTGRES_USER", "what")
	if v != "postgres" {
		t.Errorf("Expected postgres usert to be postgres, got %v instead", v)
	}
}
