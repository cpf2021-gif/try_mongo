package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"try_mongo/global"
	"try_mongo/model"
	mongodb "try_mongo/mongo"
	"try_mongo/setup"
)

func saveMdFile(dirPath string, db *mongodb.MongoDb) {
	filesInfo, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range filesInfo {
		if file.IsDir() {
			continue
		}
		// *.md
		if file.Name()[len(file.Name())-3:] != ".md" {
			continue
		}

		filename := dirPath + "/" + file.Name()
		mdContent, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		newmdfile := model.MdData{
			Title:   file.Name(),
			Content: string(mdContent),
		}

		// have old data, update
		var oldmdfile model.MdData
		err = db.FindOne(bson.D{{Key: "title", Value: file.Name()}}, &oldmdfile)
		if err == nil {
			// update
			if oldmdfile.Content != newmdfile.Content {
				err = db.ReplaceOne(bson.D{{Key: "title", Value: file.Name()}}, newmdfile)
				if err != nil {
					panic(err)
				}
			}
		} else {
			err = db.AddOne(newmdfile)
			if err != nil {
				panic(err)
			}
		}
	}
}

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

	// store md file in mongodb
	// 1. get dir
	var dirname string = "./data"

	// 2.1 create collection
	/*
			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			collection, err := client.Database("bookdb").CreateCollection(ctx, "mdfiles")
		 	if err != nil {
		 		panic(err)
		 	}
	*/

	// 2.2 get collection
	mdb := mongodb.NewMongoDb("bookdb")
	mdb.SetCollection("mdfiles")

	// 3. add md file
	saveMdFile(dirname, mdb)

	// 4. get md file
	var mdData model.MdData
	filter := bson.D{{Key: "title", Value: "first.md"}}
	err = mdb.FindOne(filter, &mdData)

	if err != nil {
		panic(err)
	}
	fmt.Println(mdData.Content)
}
