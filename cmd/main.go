package main

import (
	"os"

	"github.com/arganaphang/openmusic/internal/repository"
	"github.com/arganaphang/openmusic/internal/route"
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		zap.L().Fatal("failed to connect to database", zap.Error(err))
	}

	repositories := repository.NewRepositories(db)

	services := service.NewServices(repositories)

	app := gin.Default()
	app.SetTrustedProxies(nil)

	app.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	route.New(app, services)

	app.Run("0.0.0.0:8000")
}
