package models

import (
	"fmt"
	"github.com/google/flatbuffers/tests/namespace_test/NamespaceA"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id		bson.ObjectId	`json : "id" bson : "id"`
	Name	string			`json : "name" bson : "name"`
	Gender	string			`json : "gender" bson : "gender"`
	Age		int				`json : "age" bson : "age"`
}