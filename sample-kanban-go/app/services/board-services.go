package services

import "github.com/doug-martin/goqu/v9"

type BoardService struct {
	db *goqu.Database
}

// NewBoardService - provision service-specific code
func NewBoardService(db *goqu.Database) (*BoardService, error) {
	service := &BoardService{
		db: db,
	}
	return service, nil
}
