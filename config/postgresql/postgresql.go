package postgresql

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/NYTimes/gizmo/config"
)

// Config holds everything you need to
// connect and interact with a PostgreSQL DB.
type Config struct {
	User    string `envconfig:"POSTGRESQL_USER"`
	Pw      string `envconfig:"POSTGRESQL_PW"`
	Host    string `envconfig:"POSTGRESQL_HOST_NAME"`
	Port    int    `envconfig:"POSTGRESQL_PORT"`
	DBName  string `envconfig:"POSTGRESQL_DB_NAME"`
	SSLMode string `envconfig:"POSTGRESQL_SSL_MODE"`
}

const (
	// DefaultSSLMode is verify-full
	DefaultSSLMode = "verify-full"
	// DefaultPort is the default post for Postgresql connections
	DefaultPort = 5432
)

// DB will open a sql connection.
// Users must import a postgresql driver in their
// main to use this.
func (p *Config) DB() (*sql.DB, error) {
	db, err := sql.Open("postgres", p.String())
	if err != nil {
		return db, err
	}
	return db, nil
}

// String will return the Postgresql connection string
func (p *Config) String() string {
	if p.Port == 0 {
		p.Port = DefaultPort
	}

	if p.SSLMode != "" {
		p.SSLMode = url.QueryEscape(p.SSLMode)
	} else {
		p.SSLMode = url.QueryEscape(DefaultSSLMode)
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.User,
		p.Pw,
		p.Host,
		p.Port,
		p.DBName,
		p.SSLMode,
	)
}

// LoadConfigFromEnv will attempt to load a Postgresql object
// from environment variables. If not populated, nil
// is returned
func LoadConfigFromEnv() *Config {
	var postgres Config
	config.LoadEnvConfig(&postgres)
	if postgres.Host != "" {
		return &postgres
	}
	return nil
}
