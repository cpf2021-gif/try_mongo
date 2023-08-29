package mongo

import (
	"context"
	"time"
	"try_mongo/global"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDb struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func NewMongoDb(databaseName string) *MongoDb {
	return &MongoDb{
		Db: global.MongoClient.Database(databaseName),
	}
}

func (m *MongoDb) SetCollection(collectionName string) {
	m.Collection = m.Db.Collection(collectionName)
	if m.Collection == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := m.Db.CreateCollection(ctx, collectionName)
		if err != nil {
			panic(err)
		}
		m.Collection = m.Db.Collection(collectionName)
	}
}

func (m *MongoDb) AddOne(data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.InsertOne(ctx, data)
	return err
}

func (m *MongoDb) FindOne(filter, data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Collection.FindOne(ctx, filter).Decode(data)
}

func (m *MongoDb) ReplaceOne(filter, data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.ReplaceOne(ctx, filter, data)
	return err
}

func (m *MongoDb) DeleteOne(filter any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.DeleteOne(ctx, filter)
	return err
}
