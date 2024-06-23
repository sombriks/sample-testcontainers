package services

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/models"
)

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

func (s *BoardService) ListPerson(q string) (*[]models.Person, error) {
	var people []models.Person
	var err = s.db.From("kanban.person").
		Where(goqu.Ex{"name": goqu.Op{"ilike": fmt.Sprint("%", q, "%")}}).
		ScanStructs(&people)
	return &people, err
}
