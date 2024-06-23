package models

import "time"

type Message struct {
	Id       int64     `db:"id"`
	PersonId int64     `db:"person_id"`
	TaskId   int64     `db:"task_id"`
	Content  string    `db:"content"`
	Created  time.Time `db:"created" goqu:"omitnil"`
}
