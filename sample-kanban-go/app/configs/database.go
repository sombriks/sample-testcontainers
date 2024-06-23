package configs

import (
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
	"log"
)

// NewGoquDb - provision a query builder instance
func NewGoquDb(d *DbProps, dsn *string) (*goqu.Database, error) {
	var err error
	if d == nil {
		log.Println("[WARN] db props missing, creating a default one...")
		d, err = NewDbProps()
	}

	// configure the query builder
	if dsn == nil {
		newDsn := fmt.Sprintf("postgresql://%s:%s@%s:5432/%s?sslmode=%s", //
			d.Username, d.Password, d.Hostname, d.Database, d.SslMode)
		dsn = &newDsn
	} else {
		log.Printf("[INFO] using provided dsn [%s]\n", *dsn)
	}
	con, err := sql.Open("postgres", *dsn)
	if err != nil {
		return nil, err
	}
	// https://doug-martin.github.io/goqu/docs/selecting.html#scan-struct
	goqu.SetIgnoreUntaggedFields(true)
	db := goqu.New("postgres", con)
	db.Logger(log.Default())

	return db, nil
}
