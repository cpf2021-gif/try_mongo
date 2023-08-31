package main

import (
	"context"
	"net/http"
	"time"

	"try_mongo/adapter"
	"try_mongo/global"
	mongodb "try_mongo/mongo"
	"try_mongo/setup"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new client
	err := setup.InitializeDB()
	if err != nil {
		panic(err)
	}

	// disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func() {
		if err = global.MongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	dataSource := mongodb.NewDataSource("bookdb")

	mdController := adapter.NewMdDataController(dataSource)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routerGroupdata := e.Group("/data")
	routerGroupdata.GET("/:filename", mdController.GetMdData)
	routerGroupdata.GET("/list", mdController.GetMdDatas)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
