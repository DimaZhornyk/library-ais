package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	db *sql.DB
	r  *gin.Engine
}

func NewServer() (*Server, error) {
	psqlString := getPsqlString()
	db, err := sql.Open("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	s := &Server{}
	s.db = db
	s.r = s.initRouter()

	return s, nil
}

func (s *Server) Run() error {
	return s.r.Run(":9000")
}

func (s *Server) initRouter() *gin.Engine {
	r := gin.Default()

	h := Handler{db: s.db}
	r.GET("/queries", h.getQueries)
	r.POST("/query", h.postQuery)

	return r
}

func getPsqlString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
