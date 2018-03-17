// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ListNodePoolsOKBody list node pools o k body
// swagger:model listNodePoolsOKBody
type ListNodePoolsOKBody struct {

	// items
	Items ListNodePoolsOKBodyItems `json:"items"`
}

// Validate validates this list node pools o k body
func (m *ListNodePoolsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ListNodePoolsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListNodePoolsOKBody) UnmarshalBinary(b []byte) error {
	var res ListNodePoolsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}