package models

import "time"

type Status struct {
	Id            int64     `db:"id"`
	Description   string    `db:"name"`
	MeansComplete bool      `db:"means_complete"`
	Created       time.Time `db:"created" goqu:"omitnil"`
}
