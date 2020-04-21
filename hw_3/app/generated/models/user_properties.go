// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserProperties User info
// swagger:model UserProperties
type UserProperties struct {

	// User email address
	// Required: true
	// Min Length: 1
	Email *string `json:"email"`

	// User first name
	// Required: true
	FirstName *string `json:"firstName"`

	// User last name
	// Required: true
	LastName *string `json:"lastName"`

	// User password hash
	// Required: true
	PasswordHash *string `json:"passwordHash"`
}

// Validate validates this user properties
func (m *UserProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserProperties) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.MinLength("email", "body", string(*m.Email), 1); err != nil {
		return err
	}

	return nil
}

func (m *UserProperties) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *UserProperties) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *UserProperties) validatePasswordHash(formats strfmt.Registry) error {

	if err := validate.Required("passwordHash", "body", m.PasswordHash); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProperties) UnmarshalBinary(b []byte) error {
	var res UserProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
