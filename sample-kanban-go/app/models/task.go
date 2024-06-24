package models

import "time"

type Task struct {
	Id          int64      `db:"id" goqu:"skipinsert"`
	StatusId    int64      `db:"status_id"`
	Description string     `db:"description"`
	Created     *time.Time `db:"created" goqu:"omitnil"`
	Status      *Status
	Messages    *[]Message
	people      *[]Person
}
