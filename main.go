package main

import (
	"go-decorator/router"
	"go-decorator/database"
	"go-decorator/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()

	// database
	database.Connect()

	defer database.Destroy()

	// application middleware
	middleware.RegisterMiddleware(engine)

	// application router
	router.RegisterRouter(engine)

	// run
	engine.Run(":7001") // listen and serve on 0.0.0.0:7001 (for windows "localhost:7001")
}
