package dao

import "try_mongo/model"

type MdDataDao interface {
	FindOne(mddata *model.MdData) error
	FindMany(mddata *model.MdData) ([]model.MdData, error)
	AddOne(mddata *model.MdData) error
	UpdateOne(mddata *model.MdData) error
	DeleteOne(mddata *model.MdData) error
}
