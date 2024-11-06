package gen

import (
	"encoding/json"
	"fmt"
)

type OapiSpec struct {
	OpenAPI    string           `json:"openapi"`
	Paths      PathObject       `json:"paths"`
	Components ComponentsObject `json:"components"`
}

type PathObject map[string]PathItemObject
type PathItemObject map[string]OperationObject

type ComponentsObject struct {
	Schemas map[string]Schema `json:"schemas"`
}

type Schema struct {
	Type SchemaType `json:"type"`
	// Title      string     `json:"title"` // might add this later to generate documentation
	Enum       []string   `json:"enum"`
	Properties Properties `json:"properties"`
}

type SchemaType string

var (
	SchemaTypeArray   SchemaType = "array"
	SchemaTypeBoolean SchemaType = "boolean"
	SchemaTypeInteger SchemaType = "integer"
	SchemaTypeNull    SchemaType = "null"
	SchemaTypeNumber  SchemaType = "number"
	SchemaTypeObject  SchemaType = "object"
	SchemaTypeString  SchemaType = "string"
)

func (t *SchemaType) UnmarshalJSON(data []byte) error {
	var typ string
	if err := json.Unmarshal(data, &typ); err != nil {
		return err
	}

	switch typ {
	case "array":
		*t = SchemaTypeArray
	case "boolean":
		*t = SchemaTypeBoolean
	case "integer":
		*t = SchemaTypeInteger
	case "null":
		*t = SchemaTypeNull
	case "number":
		*t = SchemaTypeNumber
	case "object":
		*t = SchemaTypeObject
	case "string":
		*t = SchemaTypeString
	default:
		return fmt.Errorf("invalid type: %s", typ)
	}

	return nil
}

type Properties []Property

func (t *Properties) UnmarshalJSON(data []byte) error {
	var tmp map[string]*Property
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	res := make([]Property, 0)
	for k, v := range tmp {
		v.Name = k
		res = append(res, *v)
	}
	*t = res

	return nil
}

type Property struct {
	Name string
	Type string `json:"type"`
	// Title string     `json:"title"` // might add this later to generate documentation
	AnyOf []Property `json:"anyOf"`
}

type OperationObject struct {
	OperationID *string                   `json:"operationId"`
	Summary     *string                   `json:"summary,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Parameters  []ParamObject             `json:"parameters"`
	Responses   map[string]ResponseObject `json:"responses"`
}

type ParamObject struct {
	Name            string                 `json:"name"`
	Location        string                 `json:"in"`
	Description     *string                `json:"description"`
	Required        bool                   `json:"required"`
	Deprecated      bool                   `json:"deprecated"`
	AllowEmptyValue bool                   `json:"allowEmptyValue"`
	Schema          map[string]interface{} `json:"schema"`
}

type ResponseObject struct {
	Description *string                    `json:"description"`
	Content     map[string]ResponseContent `json:"content"`
}

type ResponseContent struct {
	Schema map[string]interface{} `json:"schema"`
}
