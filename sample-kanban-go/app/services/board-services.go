package services

import (
	"errors"
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

func (s *BoardService) ListPeople(q string) (*[]models.Person, error) {
	var people []models.Person
	var err = s.db.From("kanban.person").
		Where(goqu.Ex{"name": goqu.Op{"ilike": fmt.Sprint("%", q, "%")}}).
		ScanStructs(&people)
	return &people, err
}

func (s *BoardService) FindPerson(id int64) (*models.Person, error) {
	var person models.Person
	var ok, err = s.db.From("kanban.person").
		Where(goqu.C("id").Eq(id)).
		ScanStruct(&person)
	if !ok {
		return nil, errors.New(fmt.Sprint("Person #", id, " not found"))
	}
	return &person, err
}

func (s *BoardService) ListStatus() (*[]models.Status, error) {
	var statuses []models.Status
	err := s.db.From("kanban.status").
		Order(goqu.C("id").Asc()).
		ScanStructs(&statuses)
	return &statuses, err
}

func (s *BoardService) FindStatus(id int64) (*models.Status, error) {
	var status models.Status
	ok, err := s.db.From("kanban.status").
		Where(goqu.C("id").Eq(id)).
		ScanStruct(&status)
	if !ok {
		return nil, errors.New(fmt.Sprint("Status #", id, " not found"))
	}
	return &status, err
}
func (s *BoardService) FindStatusByTaskId(id int64) (*models.Status, error) {
	var status models.Status
	ok, err := s.db.From("kanban.status").
		Where(goqu.C("id").Eq(s.db.From("kanban.task").
			Select(goqu.C("status_id")).
			Where(goqu.C("id").Eq(id)))).
		ScanStruct(&status)
	if !ok {
		return nil, errors.New(fmt.Sprint("status for task #", id, " not found"))
	}
	return &status, err
}

func (s *BoardService) ListTasks(q string) (*[]models.Task, error) {
	var tasks []models.Task
	var err = s.db.From("kanban.task").
		Where(goqu.C("description").ILike(fmt.Sprint("%", q, "%"))).
		ScanStructs(&tasks)
	return &tasks, err
}

func (s *BoardService) FindTask(id int64) (*models.Task, error) {
	var task models.Task
	var ok, err = s.db.From("kanban.task").
		Where(goqu.C("id").Eq(id)).
		ScanStruct(&task)
	if !ok {
		return nil, errors.New(fmt.Sprint("Task #", id, " not found"))
	}
	return &task, err
}

func (s *BoardService) InsertTask(task *models.Task) (*models.Task, error) {
	ok, err := s.db.Insert("kanban.task").
		Rows(task).
		Returning("*").
		Executor().
		ScanStruct(task)
	if !ok {
		return nil, errors.New("failed to insert task")
	}
	return task, err
}

func (s *BoardService) UpdateTask(id int64, task *models.Task) (*models.Task, error) {
	task.Id = id
	_, err := s.db.Update("kanban.task").
		Set(*task).Where(goqu.C("id").Eq(id)).Executor().Exec()
	return task, err
}

func (s *BoardService) DeleteTask(id int64) (int64, error) {
	result, err := s.db.Delete("kanban.task").
		Where(goqu.C("id").Eq(id)).
		Executor().
		Exec()
	affected, err := result.RowsAffected()
	return affected, err
}
