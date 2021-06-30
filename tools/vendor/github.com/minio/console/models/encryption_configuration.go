// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2021 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncryptionConfiguration encryption configuration
//
// swagger:model encryptionConfiguration
type EncryptionConfiguration struct {
	MetadataFields

	// aws
	Aws *AwsConfiguration `json:"aws,omitempty"`

	// client
	Client *KeyPairConfiguration `json:"client,omitempty"`

	// gcp
	Gcp *GcpConfiguration `json:"gcp,omitempty"`

	// gemalto
	Gemalto *GemaltoConfiguration `json:"gemalto,omitempty"`

	// image
	Image string `json:"image,omitempty"`

	// replicas
	Replicas string `json:"replicas,omitempty"`

	// server
	Server *KeyPairConfiguration `json:"server,omitempty"`

	// vault
	Vault *VaultConfiguration `json:"vault,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *EncryptionConfiguration) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MetadataFields
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MetadataFields = aO0

	// AO1
	var dataAO1 struct {
		Aws *AwsConfiguration `json:"aws,omitempty"`

		Client *KeyPairConfiguration `json:"client,omitempty"`

		Gcp *GcpConfiguration `json:"gcp,omitempty"`

		Gemalto *GemaltoConfiguration `json:"gemalto,omitempty"`

		Image string `json:"image,omitempty"`

		Replicas string `json:"replicas,omitempty"`

		Server *KeyPairConfiguration `json:"server,omitempty"`

		Vault *VaultConfiguration `json:"vault,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Aws = dataAO1.Aws

	m.Client = dataAO1.Client

	m.Gcp = dataAO1.Gcp

	m.Gemalto = dataAO1.Gemalto

	m.Image = dataAO1.Image

	m.Replicas = dataAO1.Replicas

	m.Server = dataAO1.Server

	m.Vault = dataAO1.Vault

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m EncryptionConfiguration) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MetadataFields)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Aws *AwsConfiguration `json:"aws,omitempty"`

		Client *KeyPairConfiguration `json:"client,omitempty"`

		Gcp *GcpConfiguration `json:"gcp,omitempty"`

		Gemalto *GemaltoConfiguration `json:"gemalto,omitempty"`

		Image string `json:"image,omitempty"`

		Replicas string `json:"replicas,omitempty"`

		Server *KeyPairConfiguration `json:"server,omitempty"`

		Vault *VaultConfiguration `json:"vault,omitempty"`
	}

	dataAO1.Aws = m.Aws

	dataAO1.Client = m.Client

	dataAO1.Gcp = m.Gcp

	dataAO1.Gemalto = m.Gemalto

	dataAO1.Image = m.Image

	dataAO1.Replicas = m.Replicas

	dataAO1.Server = m.Server

	dataAO1.Vault = m.Vault

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this encryption configuration
func (m *EncryptionConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetadataFields
	if err := m.MetadataFields.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGemalto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVault(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionConfiguration) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) validateClient(formats strfmt.Registry) error {

	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) validateGemalto(formats strfmt.Registry) error {

	if swag.IsZero(m.Gemalto) { // not required
		return nil
	}

	if m.Gemalto != nil {
		if err := m.Gemalto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gemalto")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) validateServer(formats strfmt.Registry) error {

	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) validateVault(formats strfmt.Registry) error {

	if swag.IsZero(m.Vault) { // not required
		return nil
	}

	if m.Vault != nil {
		if err := m.Vault.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this encryption configuration based on the context it is used
func (m *EncryptionConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MetadataFields
	if err := m.MetadataFields.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGemalto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVault(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionConfiguration) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {
		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {
		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) contextValidateGemalto(ctx context.Context, formats strfmt.Registry) error {

	if m.Gemalto != nil {
		if err := m.Gemalto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gemalto")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *EncryptionConfiguration) contextValidateVault(ctx context.Context, formats strfmt.Registry) error {

	if m.Vault != nil {
		if err := m.Vault.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vault")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionConfiguration) UnmarshalBinary(b []byte) error {
	var res EncryptionConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
