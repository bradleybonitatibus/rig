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

package pg

import (
	"testing"
)

const pg = "postgres"

func TestConnectionConfig_ODBC(t *testing.T) {
	type test struct {
		name string
		cfg  *ConnectionConfig
		want string
	}

	tests := []test{
		{
			name: "all fields given",
			cfg: &ConnectionConfig{
				Host:     pg,
				User:     pg,
				Password: pg,
				Database: pg,
				SSLMode:  "verify-ca",
				Port:     5432,
			},
			want: "host=postgres port=5432 user=postgres dbname=postgres password=postgres sslmode=verify-ca",
		},
		{
			name: "sslmode is prefer and shouldn't appear in ODBC",
			cfg: &ConnectionConfig{
				Host:     pg,
				User:     pg,
				Password: pg,
				Database: pg,
				SSLMode:  "prefer",
				Port:     5432,
			},
			want: "host=postgres port=5432 user=postgres dbname=postgres password=postgres",
		},
		{
			name: "omit port to default use 5432",
			cfg: &ConnectionConfig{
				Host:     pg,
				User:     pg,
				Password: pg,
				Database: pg,
				SSLMode:  "prefer",
			},
			want: "host=postgres port=5432 user=postgres dbname=postgres password=postgres",
		},
	}

	t.Parallel()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.cfg.ODBC()
			if tc.want != got {
				t.Errorf("%v: expected %v, got %v instead", tc.name, tc.want, got)
			}
		})
	}
}
