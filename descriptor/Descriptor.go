package descriptor

import "gopkg.in/mgo.v2/bson"

type Descriptor struct {
	Id          bson.ObjectId `bson:"_id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Example     string `bson:"example" json:"example"`
	Url         string `bson:"url" json:"url"`
	Tags        []string `bson:"tags" json:"tags"`
}
