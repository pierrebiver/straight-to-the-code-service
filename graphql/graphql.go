package graphql

import (
	"github.com/neelance/graphql-go"
	"github.com/straight-to-the-code-service/model"
	"github.com/straight-to-the-code-service/mongo"
	"fmt"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
		descriptors(): [Descriptor]
	}

	type Mutation {
		add(descriptor: DescriptorInput!): Descriptor
	}

	input DescriptorInput {
		name: String!
		description: String
		tags: [String!]!
	}

	type Descriptor {
		id: ID!
		name: String!
		description: String
		tags: [String!]!
	}
`

type Resolver struct{}

type descriptorResolver struct {
	Descriptor model.Descriptor
}

type DescriptorArgs struct {
	Descriptor *model.DescriptorInput
}

func (d *descriptorResolver) ID() graphql.ID {
	return d.Descriptor.ID
}

func (d *descriptorResolver) Name() string {
	return d.Descriptor.Name
}

func (d *descriptorResolver) Description() *string {
	return &d.Descriptor.Description
}

func (d *descriptorResolver) Tags() []string {
	return d.Descriptor.Tags
}

func (r *Resolver) Descriptors() *[]*descriptorResolver {
	descriptors, _ := mongo.Descriptors()

	descriptorResolvers := make([]*descriptorResolver, 0)
	for _, d := range descriptors {
		descriptorResolvers = append(descriptorResolvers, &descriptorResolver{d})
	}

	return &descriptorResolvers
}

func (r *Resolver) Add(args DescriptorArgs) *descriptorResolver {
	input := args.Descriptor
	mongo.Add(input)
	fmt.Print(input)

	return &descriptorResolver{model.Descriptor{ID: input.ID, Name: input.Name, Description: *input.Description, Tags: input.Tags}}
}
