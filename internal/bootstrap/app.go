package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/vector-ops/lifefolio/configs"
	"github.com/vector-ops/lifefolio/internal/bootstrap/database"
	"github.com/vector-ops/lifefolio/internal/bootstrap/web"
)

type App struct {
	server *echo.Echo
	db     *database.Database
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	a.setupEnv()
	a.setupHTTP()
	a.setupDatabase()

	port := fmt.Sprintf(":%s", configs.GetEnv("PORT", "3000"))
	err := a.server.Start(port)
	if err != nil {
		log.Println(err)
		a.closeServices()
		os.Exit(1)
	}
}

func (a *App) setupEnv() {
	configs.LoadEnv()
}

func (a *App) setupHTTP() {
	a.server = web.SetupServer()

	middlewares := web.NewMiddleware(a.server, a.db)
	middlewares.Init()
}

func (a *App) setupDatabase() {
	db := database.NewDatabase()
	db.SetupEntClient()
	db.SetupRedis()

	a.db = db
}

func (a *App) closeServices() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Println("Failed to shutdown server.")
	}

	if err := a.db.EntClient.Close(); err != nil {
		log.Println("Failed to shutdown Postgres database")
	}

	if err := a.db.RDB.Shutdown(ctx); err != nil {
		log.Println("Failed to shutdown Redis database")
	}

}
