package models

import "time"

type Task struct {
	Id          int64     `db:"id"`
	StatusId    int64     `db:"status_id"`
	Description string    `db:"description"`
	Created     time.Time `db:"created" goqu:"omitnil"`
}
