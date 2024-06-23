package models

import "time"

type Person struct {
	Id      int64     `db:"id"`
	Name    string    `db:"name"`
	Created time.Time `db:"created" goqu:"omitnil"`
}
