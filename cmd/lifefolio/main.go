package main

import (
	"github.com/vector-ops/lifefolio/internal/bootstrap"
)

func main() {
	app := bootstrap.NewApp()
	app.Run()
}
