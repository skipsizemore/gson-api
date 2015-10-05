package gsonapi

import (
	"encoding/json"

	gas "github.com/obieq/gas"
	validations "github.com/obieq/goar-validations"
)

type JsonApiResourcer interface {
	Resourcer
}

type JsonApiResource struct {
	Data   interface{}    `json:"data"`
	Errors []JsonApiError `json:"errors,omitempty"`
	Meta   interface{}    `json:"meta,omitempty"`
}

type JsonApiPayload struct {
	Body interface{}
}

type Resourcer interface {
	// URL() string
	// URI() string
	GetID() string
	SetID(string) error
	// BuildLinks()
	MapToModel(model interface{}) error
	MapFromModel(model interface{}) error
	Errors() []JsonApiError
	SetErrors(map[string]*validations.ValidationError)
	// Atts() interface{}
	// SetAtts(interface{})
}

type Resource struct {
	// ResourceType string      `json:"type,omitempty"`
	// ID           string      `json:"id,omitempty"`
	ID string `json:"id,omitempty" jsonapi:"-"`
	// ID string `json:"id,omitempty" jsonapi:"name=id"`
	// Attributes   interface{} `json:"attributes,omitempty"`
	errors map[string]*validations.ValidationError
}

type Link struct {
	Self    string `json:"self,omitempty"`
	Related string `json:"related,omitempty"`
}

// CollectionLink => JSON API links
type CollectionLink struct {
	Self    string    `json:"self,omitempty"`
	Related string    `json:"related,omitempty"`
	Linkage []Linkage `json:"linkage,omitempty"`
}

// Linkage => JSON API linkage
type Linkage struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type JsonApiErrorLink struct {
	About string `json:"about,omitempty"`
}

type JsonApiErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

type JsonApiError struct {
	ID     string              `json:"id,omitempty"`
	Status string              `json:"status,omitempty"`
	Code   string              `json:"code,omitempty"`
	Title  string              `json:"title,omitempty"`
	Detail string              `json:"detail,omitempty"`
	Links  *JsonApiErrorLink   `json:"linkz,omitempty"`
	Source *JsonApiErrorSource `json:"source,omitempty"`
}

func (r Resource) GetID() string {
	return r.ID
}

func (r *Resource) SetID(id string) error {
	r.ID = id
	return nil
}

// func (r *Resource) URL() string {
// 	return Config.URL
// }
//
// func LinkSelfCollection(r Resourcer) string {
// 	return r.URL() + r.URI()
// }
//
// func LinkSelfInstance(r Resourcer) string {
// 	return LinkSelfCollection(r) + "/" + r.GetID()
// }

func (r *Resource) Errors() []JsonApiError {
	errors := []JsonApiError{}
	var err JsonApiError

	for k, v := range r.errors {
		key := gas.String(k).Dasherize()
		err = JsonApiError{Detail: v.Message, Status: "422"}
		err.Source = &JsonApiErrorSource{Pointer: "data/attributes/" + key}
		errors = append(errors, err)
	}

	return errors
}

func (r *Resource) SetErrors(errors map[string]*validations.ValidationError) {
	r.errors = errors
}

func UnmarshalJsonApiData(source interface{}, destination interface{}) error {
	var err error

	if tmp, err := json.Marshal(source); err == nil {
		err = json.Unmarshal(tmp, &destination)
	}
	return err
}
