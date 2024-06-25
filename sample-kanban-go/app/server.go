package app

import (
	"embed"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/configs"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/requests"
	"github.com/sombriks/sample-testcontainers/sample-kanban-go/app/services"
	"log"
	"net/http"
	"os"
)

type KanbanServer struct {
	controller *requests.BoardRequest
	service    *services.BoardService
	db         *goqu.Database
	e          *echo.Echo
}

//go:embed static
var staticFS embed.FS

func NewKanbanServer(db *goqu.Database) (*KanbanServer, error) {
	// configuration phase
	var err error

	if db == nil {
		log.Println("[WARN] db is nil, provisioning a default one")
		db, err = configs.NewGoquDb(nil, nil)
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

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(staticFS),
		Root:       "static",
		//Browse:     true,
	}))

	// routes/requests
	e.GET("/", controller.Index)

	e.GET("/board", controller.BoardPage, controller.CookieCheck)

	login := e.Group("/login")
	login.GET("", controller.LoginPage)
	login.POST("", controller.FakeLogin)

	e.GET("/logout", controller.FakeLogout)

	e.GET("/table", controller.TablePage, controller.CookieCheck)

	task := e.Group("/task", controller.CookieCheck)
	task.POST("", controller.AddTask)

	taskId := task.Group("/:id")
	taskId.PUT("", controller.UpdateTask)
	taskId.DELETE("", controller.DeleteTask)
	taskId.DELETE("/person/:personId", controller.RemovePerson)
	taskId.POST("/join", controller.JoinTask)
	taskId.POST("/comments", controller.AddComment)

	server := &KanbanServer{
		controller: controller,
		service:    service,
		db:         db,
		e:          e,
	}

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
