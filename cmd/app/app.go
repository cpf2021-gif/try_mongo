package main

import (
	"context"
	"net/http"
	"time"

	"try_mongo/global"
	"try_mongo/model"
	mongodb "try_mongo/mongo"
	"try_mongo/setup"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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

	// get collection
	mdb := mongodb.NewMongoDb("bookdb")
	mdb.SetCollection("mdfiles")

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routerGroupdata := e.Group("/data")
	routerGroupdata.GET("/:filename", func(c echo.Context) error {
		filename := c.Param("filename")
		var mdfile model.MdData
		err := mdb.FindOne(bson.D{{Key: "title", Value: filename}}, &mdfile)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "not found",
			})
		}
		return c.JSON(http.StatusOK, mdfile)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
