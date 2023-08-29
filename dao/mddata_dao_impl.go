package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"try_mongo/model"
)

type MdDataDaoImpl struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

func (m *MdDataDaoImpl) FindOne(mddata *model.MdData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return m.Collection.FindOne(ctx, bson.D{{Key: "title", Value: mddata.Title}}).Decode(mddata)
}

func (m *MdDataDaoImpl) AddOne(mddata *model.MdData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.InsertOne(ctx, mddata)
	return err
}

func (m *MdDataDaoImpl) UpdateOne(mddata *model.MdData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.ReplaceOne(ctx, bson.D{{Key: "title", Value: mddata.Title}}, mddata)
	return err
}

func (m *MdDataDaoImpl) DeleteOne(mddate *model.MdData) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.Collection.DeleteOne(ctx, bson.D{{Key: "title", Value: mddate.Title}})
	return err
}
