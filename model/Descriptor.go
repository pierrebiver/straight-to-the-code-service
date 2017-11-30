package model

import (
	"github.com/neelance/graphql-go"
)

type Descriptor struct {
	ID          graphql.ID `bson:"_id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	Example     string     `bson:"example" json:"example"`
	Tags        []string   `bson:"tags" json:"tags"`
}
