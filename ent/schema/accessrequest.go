package schema

import "entgo.io/ent"

// AccessRequest holds the schema definition for the AccessRequest entity.
type AccessRequest struct {
	ent.Schema
}

// Fields of the AccessRequest.
func (AccessRequest) Fields() []ent.Field {
	return nil
}

// Edges of the AccessRequest.
func (AccessRequest) Edges() []ent.Edge {
	return nil
}
