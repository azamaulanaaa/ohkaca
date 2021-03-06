package schema

import (
	"errors"

	gql "github.com/graphql-go/graphql"
)

type Schema struct {
	query    gql.Fields
	mutation gql.Fields
}

func New() *Schema {
	return &Schema{
		query:    make(gql.Fields),
		mutation: make(gql.Fields),
	}
}

func (s *Schema) AddQuery(key string, field *gql.Field) error {
	if _, exist := s.query[key]; exist {
		return errors.New("Query key already used")
	}

	s.query[key] = field
	return nil
}

func (s *Schema) AddMutation(key string, field *gql.Field) error {
	if _, exist := s.mutation[key]; exist {
		return errors.New("Mutation key already used")
	}

	s.mutation[key] = field
	return nil
}

func (s *Schema) Generate() (gql.Schema, error) {
	return gql.NewSchema(gql.SchemaConfig{
		Query:    genGQLObjectFromFields("Query", s.query),
		Mutation: genGQLObjectFromFields("Mutation", s.mutation),
	})
}

func genGQLObjectFromFields(name string, fields gql.Fields) *gql.Object {
	if len(fields) > 0 {
		return gql.NewObject(gql.ObjectConfig{
			Name:   name,
			Fields: fields,
		})
	}

	return nil
}
