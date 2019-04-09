package avro

import (
	"encoding/json"

	"github.com/valyala/fastjson"
)

// UnionSchema - A JSON array, representing a union of embedded types.
type UnionSchema []Schema

// TypeName -
func (t UnionSchema) TypeName() Type {
	return TypeUnion
}

// MarshalJSON -
func (t UnionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(t)
}

func translateValues2UnionSchema(values []*fastjson.Value) (Schema, error) {
	union := UnionSchema(make([]Schema, 0, len(values)))
	for _, value := range values {
		schema, err := translateValue2AnySchema(value)
		if err != nil {
			return nil, err
		}
		union = append(union, schema)
	}
	return union, nil
}
