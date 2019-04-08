// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListOryAccessControlPoliciesOK ListOryAccessControlPoliciesOK ListOryAccessControlPoliciesOK ListOryAccessControlPoliciesOK handles this case with default header values.
//
// Policies is an array of policies.
// swagger:model ListOryAccessControlPoliciesOK
type ListOryAccessControlPoliciesOK struct {

	// payload
	Payload []*Policy `json:"Payload"`
}

// Validate validates this list ory access control policies o k
func (m *ListOryAccessControlPoliciesOK) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListOryAccessControlPoliciesOK) validatePayload(formats strfmt.Registry) error {

	if swag.IsZero(m.Payload) { // not required
		return nil
	}

	for i := 0; i < len(m.Payload); i++ {
		if swag.IsZero(m.Payload[i]) { // not required
			continue
		}

		if m.Payload[i] != nil {
			if err := m.Payload[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Payload" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListOryAccessControlPoliciesOK) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListOryAccessControlPoliciesOK) UnmarshalBinary(b []byte) error {
	var res ListOryAccessControlPoliciesOK
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}