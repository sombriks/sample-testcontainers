package app

import (
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/configs"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/requests"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/services"
	"log"
	"os"
)

type KanbanServer struct {
	controller *requests.BoardRequest
	service    *services.BoardService
	db         *goqu.Database
	e          *echo.Echo
}

func NewKanbanServer(db *goqu.Database) (*KanbanServer, error) {
	var err error

	if db == nil {
		log.Println("db is nil, provisioning a default one")
		db, err = configs.NewGoquDb()
		if err != nil {
			return nil, err
		}
	}
	service, err := services.NewBoardService(db)
	if err != nil {
		return nil, err
	}
	controller, err := requests.NewBoardRequest(service)
	if err != nil {
		return nil, err
	}

	e := echo.New()

	// configuration phase
	server := &KanbanServer{
		controller: controller,
		service:    service,
		db:         db,
		e:          e,
	}

	// Middlewares
	server.e.Use(middleware.Logger())
	server.e.Use(middleware.Recover())

	// routes/requests
	server.e.GET("/", controller.Index)

	server.e.GET("/board", controller.BoardPage, controller.CookieCheck)

	login := server.e.Group("/login")
	login.GET("", controller.LoginPage)
	login.POST("", controller.FakeLogin)

	server.e.GET("/logout", controller.FakeLogout)

	server.e.GET("/table", controller.TablePage, controller.CookieCheck)

	task := server.e.Group("/task", controller.CookieCheck)
	task.POST("/", controller.AddTask)

	taskId := task.Group("/:id")
	taskId.PUT("/", controller.UpdateTask)
	taskId.DELETE("/", controller.DeleteTask)
	taskId.DELETE("/person/:personId", controller.DeleteTask)
	taskId.POST("/join", controller.JoinTask)
	taskId.POST("/comments", controller.AddComent)

	return server, nil
}

func (server *KanbanServer) CheckDb() error {
	_, err := server.db.Exec("select 1 + 1 as result")
	return err
}

func (server *KanbanServer) Listen() error {

	port, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("PORT environment variable not set")
	}

	return server.e.Start(fmt.Sprintf(":%s", port))
}
