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

func (m *MdDataDaoImpl) FindMany(mddata *model.MdData) ([]model.MdData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var cursor *mongo.Cursor
	var err error

	if mddata.Title == "" {
		cursor, err = m.Collection.Find(ctx, bson.D{})
	} else {
		cursor, err = m.Collection.Find(ctx, bson.D{{Key: "title", Value: mddata.Title}})
	}

	if err != nil {
		return nil, err
	}

	var mddatas []model.MdData
	for cursor.Next(ctx) {
		var mddata model.MdData
		if err := cursor.Decode(&mddata); err != nil {
			return nil, err
		}
		mddatas = append(mddatas, mddata)
	}
	return mddatas, nil
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
