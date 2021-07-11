// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Pet pet
//
// swagger:model Pet
type Pet struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Example: doggie
	// Required: true
	Name *string `json:"name"`

	// photo urls
	// Required: true
	PhotoUrls []string `json:"photoUrls" xml:"photoUrl"`

	// pet status in the store
	// Enum: [available pending sold]
	Status string `json:"status,omitempty"`
}

// Validate validates this pet
func (m *Pet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoUrls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Pet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Pet) validatePhotoUrls(formats strfmt.Registry) error {

	if err := validate.Required("photoUrls", "body", m.PhotoUrls); err != nil {
		return err
	}

	return nil
}

var petTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["available","pending","sold"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		petTypeStatusPropEnum = append(petTypeStatusPropEnum, v)
	}
}

const (

	// PetStatusAvailable captures enum value "available"
	PetStatusAvailable string = "available"

	// PetStatusPending captures enum value "pending"
	PetStatusPending string = "pending"

	// PetStatusSold captures enum value "sold"
	PetStatusSold string = "sold"
)

// prop value enum
func (m *Pet) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, petTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Pet) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pet based on context it is used
func (m *Pet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Pet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Pet) UnmarshalBinary(b []byte) error {
	var res Pet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
