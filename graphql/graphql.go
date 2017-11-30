package graphql

import (
	"github.com/straight-to-the-code-service/model"
	"github.com/neelance/graphql-go"
	"github.com/straight-to-the-code-service/mongo"
)

var Schema = `
	schema {
		query: Query
	}

	type Query {
		descriptors(): [Descriptor!]!
	}

	type Descriptor {
		id: ID!
		name: String!
		Description: String
		Example: String
		tags: [String!]!
	}
`

type Resolver struct{}

type descriptorResolver struct {
	descriptor model.Descriptor
}

func (d *descriptorResolver) ID() graphql.ID {
	return d.descriptor.ID
}

func (d *descriptorResolver) Name() string {
	return d.descriptor.Name
}

func (d *descriptorResolver) Description() *string {
	return &d.descriptor.Description
}

func (d *descriptorResolver) Example() *string {
	return &d.descriptor.Example
}

func (d *descriptorResolver) Tags() []string {
	return d.descriptor.Tags
}

func (r *Resolver) Descriptors() []descriptorResolver {
	descriptors, _ := mongo.Descriptors()

	descriptorResolvers := make([]descriptorResolver, 0)
	for _, d := range descriptors {
		descriptorResolvers = append(descriptorResolvers, descriptorResolver{d})
	}

	return descriptorResolvers
}
