package api

import (
	"fmt"
	"os"

	"github.com/cristianchaparroa/merlin/bycles/initializer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type IServer interface {
}

type Server struct {
	api *gin.Engine
	db  *gorm.DB
}

func NewServer() *Server {
	return &Server{api: gin.Default()}
}

func (s *Server) SetupDB() {
	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASSWORD_DB")
	dbName := os.Getenv("NAME_DB")
	host := os.Getenv("HOST_DB")

	datasource := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbName)

	db, err := gorm.Open("postgres", datasource)

	if err != nil {
		panic(err)
	}
	s.db = db
}

func (s *Server) SetupRoutes() {

	bh := NewBicycleHandler(s.db)

	s.api.GET("/v1/bicycles", bh.GetAll)
	s.api.GET("/v1/bicycles/:id", bh.GetById)
	s.api.POST("/v1/bicycles", bh.Create)
	s.api.PUT("/v1/bicycles/:id", bh.Update)
}

func (s *Server) Run() {
	s.SetupDB()
	initializer.InitModels(s.db)
	s.SetupRoutes()
	s.api.Run()
}
