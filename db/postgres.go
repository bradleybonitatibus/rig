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

package db

import "fmt"

// SSLMode type contains the different options to provide as an option
// to the sslmode connection string parameter.
// See more at https://www.postgresql.org/docs/13/libpq-connect.html#LIBPQ-CONNECT-SSLMODE.
type SSLMode string

const (
	SSLModeDisable    SSLMode = "disable"
	SSLModeAllow      SSLMode = "allow"
	SSLModeRequire    SSLMode = "require"
	SSLModePrefer     SSLMode = "prefer" // Prefer is the libpq default.
	SSLModeVerifyCA   SSLMode = "verify-ca"
	SSLModeVerifyFull SSLMode = "verify-full"
)

// ConnectionConfig is the configuration to connect to a Postgres Database.
// Additional to the ODBC information, there are connection pool configuration
// fields that can be set in this object.
type ConnectionConfig struct {
	Host            string  `json:"host" yaml:"host"`
	User            string  `json:"user" yaml:"user"`
	Password        string  `json:"password" yaml:"password"`
	Database        string  `json:"database" yaml:"database"`
	Port            int     `json:"port" yaml:"port"`
	SSLMode         SSLMode `json:"sslmode" yaml:"sslmode"`
	MaxOpenConns    int     `json:"max_open_conns" yaml:"max_open_conns"`
	ConnMaxLifetime int     `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
	MaxIdleConns    int     `json:"max_idle_conns" yaml:"max_idle_conns"`
	ConnMaxIdleTime int     `json:"conn_max_idle_time" yaml:"conn_max_idle_time"`
}

// ODBC formats the ConnectionConfig config into an ODBC string format.
func (c *ConnectionConfig) ODBC() string {
	if c.Port == 0 {
		c.Port = 5432
	}
	strTemplate := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", c.Host,
		c.Port,
		c.User,
		c.Database,
		c.Password,
	)

	if c.SSLMode == "" || c.SSLMode != SSLModePrefer {
		return fmt.Sprintf("%v sslmode=%s", strTemplate, c.SSLMode)
	}
	return strTemplate
}
