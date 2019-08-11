package models

import "gopkg.in/mgo.v2/bson"

type Company struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Zip         string        `bson:"zip" json:"zip"`
	Website     string        `bson:"website" json:"website"`
}

type DatabasePath struct{
	Path        string        `bson:"path json:"path"`
}

