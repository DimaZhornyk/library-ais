package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	db *sqlx.DB
	r  *gin.Engine
}

func NewServer() (*Server, error) {
	psqlString := getPsqlString()
	db, err := sqlx.Open("postgres", psqlString)
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
	if err := migrate(s.db); err != nil {
		return err
	}

	return s.r.Run(":9000")
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func (s *Server) initRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	h := Handler{db: s.db}

	r.GET("/queries", h.getQueries)
	r.POST("/query", h.postQuery)

	c := r.Group("/complex")
	c.POST("/createBookInstance", h.createBookInstance)

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
