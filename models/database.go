package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	Driver   string
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

func (d Database) toDB() (*sql.DB, error) {
	return sql.Open(d.Driver, d.connectionString())
}

func (d Database) connectionString() string {
	str := ""

	if d.Name != "" {
		str += fmt.Sprintf(" dbname=%s", d.Name)
	}

	if d.User != "" {
		str += fmt.Sprintf(" user=%s", d.User)

		if d.Password != "" {
			str += fmt.Sprintf(" password=%s", d.Password)
		}
	}

	if d.Host != "" {
		str += fmt.Sprintf(" host=%s", d.Host)
	}

	if d.Port != "" {
		str += fmt.Sprintf(" port=%s", d.Port)
	}

	if d.SSLMode != "" {
		str += fmt.Sprintf(" sslmode=%s", d.SSLMode)
	}

	return str
}
