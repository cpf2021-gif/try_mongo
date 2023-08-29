package mongo

import (
	"try_mongo/dao"
	"try_mongo/global"

	"go.mongodb.org/mongo-driver/mongo"
)

type DateSource struct {
	Db *mongo.Database
}

func NewDataSource(dbName string) *DateSource {
	return &DateSource{
		Db: global.MongoClient.Database(dbName),
	}
}

func (d *DateSource) MdDataDao() dao.MdDataDao {
	return &dao.MdDataDaoImpl{
		DB:         d.Db,
		Collection: d.Db.Collection("mdfiles"),
	}
}
