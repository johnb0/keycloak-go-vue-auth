// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TokenPair token pair
//
// swagger:model TokenPair
type TokenPair struct {

	// access token
	AccessToken AccessToken `json:"AccessToken,omitempty"`

	// refresh token
	RefreshToken RefreshToken `json:"RefreshToken,omitempty"`
}

// Validate validates this token pair
func (m *TokenPair) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenPair) validateAccessToken(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessToken) { // not required
		return nil
	}

	if err := m.AccessToken.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccessToken")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AccessToken")
		}
		return err
	}

	return nil
}

func (m *TokenPair) validateRefreshToken(formats strfmt.Registry) error {
	if swag.IsZero(m.RefreshToken) { // not required
		return nil
	}

	if err := m.RefreshToken.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RefreshToken")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("RefreshToken")
		}
		return err
	}

	return nil
}

// ContextValidate validate this token pair based on the context it is used
func (m *TokenPair) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefreshToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenPair) contextValidateAccessToken(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccessToken.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccessToken")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("AccessToken")
		}
		return err
	}

	return nil
}

func (m *TokenPair) contextValidateRefreshToken(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RefreshToken.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RefreshToken")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("RefreshToken")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenPair) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenPair) UnmarshalBinary(b []byte) error {
	var res TokenPair
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
