// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertBlock alert block
//
// swagger:model AlertBlock
type AlertBlock struct {

	// chain Id
	// Example: 1337
	ChainID int64 `json:"chainId,omitempty"`

	// hash
	// Example: 0xf9e777b739cf90a197c74c461933422dcf26fadf50e0ef9aa72af76727da87ca
	Hash string `json:"hash,omitempty"`

	// number
	// Example: 1235678901234
	Number int64 `json:"number,omitempty"`

	// Timestamp (RFC3339)
	// Example: 2022-03-01T12:24:33Z
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this alert block
func (m *AlertBlock) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alert block based on context it is used
func (m *AlertBlock) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertBlock) UnmarshalBinary(b []byte) error {
	var res AlertBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
