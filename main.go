package main

import (
	"github.com/gin-gonic/gin"

	"github.com/afdolriski/restful-test/api"
	"github.com/afdolriski/restful-test/database"
)

func main() {

	db, _ := database.InitDB()

	app := gin.Default()
	app.Use(database.Inject(db))
	api.ApplyRoutes(app)

	app.Run()
}
