package models

import "time"

type Status struct {
	Id            int64      `db:"id" goqu:"skipinsert"`
	Description   string     `db:"description"`
	MeansComplete bool       `db:"means_complete"`
	Created       *time.Time `db:"created" goqu:"omitnil"`
}
