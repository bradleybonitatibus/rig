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

import "os"

// MustGetEnv returns an environment variable, and panics if the variable
// is not set (i.e. returns an empty string).
func MustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(k + " env var is not set")
	}
	return v
}

// GetEnvWithFallback attempts to get an environment variable, and returns
// the `fallback` value if the variable was not set.
func GetEnvWithFallback(k, fallback string) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	return v
}
