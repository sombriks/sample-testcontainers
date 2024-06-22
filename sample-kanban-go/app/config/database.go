package config

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// NewGoquDb - provision a query builder instance
func NewGoquDb() (*goqu.Database, error) {

	// collect configuration from environment
	database, ok := os.LookupEnv("PG_DATABASE")
	if !ok {
		return nil, errors.New("PG_DATABASE environment variable not set")
	}
	username, ok := os.LookupEnv("PG_USERNAME")
	if !ok {
		return nil, errors.New("PG_USERNAME environment variable not set")
	}
	password, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		return nil, errors.New("PG_PASSWORD environment variable not set")
	}
	hostname, ok := os.LookupEnv("PG_HOSTNAME")
	if !ok {
		return nil, errors.New("PG_HOSTNAME environment variable not set")
	}
	sslMode, ok := os.LookupEnv("PG_SSL_MODE")
	if !ok {
		return nil, errors.New("PG_SSL_MODE environment variable not set")
	}

	// configure the query builder
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=%s", //
		username, password, hostname, database, sslMode)
	con, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	goqu := goqu.New("postgres", con)
	goqu.Logger(log.Default())

	return goqu, nil
}
