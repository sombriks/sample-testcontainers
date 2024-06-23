package services

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/joho/godotenv"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/configs"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

type ServiceTestSuit struct {
	suite.Suite
	ctx     context.Context
	tc      *postgres.PostgresContainer
	db      *goqu.Database
	service *BoardService
}

// TestRunSuite when writing suites this is needed as a 'suite entrypoint'
// see https://pkg.go.dev/github.com/stretchr/testify/suite
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuit))
}

func (s *ServiceTestSuit) SetupSuite() {
	var err error
	// Test execution point is inside the package, not in project root
	_ = godotenv.Load("../../.env")

	s.ctx = context.Background()

	props, err := configs.NewDbProps()
	if err != nil {
		s.Fail("Suite setup failed", err)
	}
	s.tc, err = postgres.RunContainer(s.ctx,
		testcontainers.WithImage("postgres:16.3-alpine3.20"),
		postgres.WithInitScripts(fmt.Sprint("../../", props.InitScript)), // path changes due test entrypoint
		postgres.WithUsername(props.Username),
		postgres.WithDatabase(props.Database),
		postgres.WithPassword(props.Password),
		testcontainers.WithWaitStrategy(wait.
			ForLog("database system is ready to accept connections").
			WithOccurrence(2).
			WithStartupTimeout(10*time.Second)))
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

	dsn, err := s.tc.ConnectionString(s.ctx, fmt.Sprint("sslmode=", props.SslMode))
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

	s.db, err = configs.NewGoquDb(nil, &dsn)
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

	s.service, err = NewBoardService(s.db)
	if err != nil {
		s.Fail("Suite setup failed", err)
	}

}

func (s *ServiceTestSuit) TearDownSuite() {
	err := s.tc.Terminate(s.ctx)
	if err != nil {
		s.Fail("Suite tear down failed", err)
	}
}

func (s *ServiceTestSuit) TestShouldListPeople() {

	people, err := s.service.ListPerson("")
	s.Nil(err)
	s.Len(*people, 5)
}
