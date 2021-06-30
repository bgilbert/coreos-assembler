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

package admin_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// GetTenantUsageOKCode is the HTTP code returned for type GetTenantUsageOK
const GetTenantUsageOKCode int = 200

/*GetTenantUsageOK A successful response.

swagger:response getTenantUsageOK
*/
type GetTenantUsageOK struct {

	/*
	  In: Body
	*/
	Payload *models.TenantUsage `json:"body,omitempty"`
}

// NewGetTenantUsageOK creates GetTenantUsageOK with default headers values
func NewGetTenantUsageOK() *GetTenantUsageOK {

	return &GetTenantUsageOK{}
}

// WithPayload adds the payload to the get tenant usage o k response
func (o *GetTenantUsageOK) WithPayload(payload *models.TenantUsage) *GetTenantUsageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tenant usage o k response
func (o *GetTenantUsageOK) SetPayload(payload *models.TenantUsage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTenantUsageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetTenantUsageDefault Generic error response.

swagger:response getTenantUsageDefault
*/
type GetTenantUsageDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTenantUsageDefault creates GetTenantUsageDefault with default headers values
func NewGetTenantUsageDefault(code int) *GetTenantUsageDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTenantUsageDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get tenant usage default response
func (o *GetTenantUsageDefault) WithStatusCode(code int) *GetTenantUsageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get tenant usage default response
func (o *GetTenantUsageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get tenant usage default response
func (o *GetTenantUsageDefault) WithPayload(payload *models.Error) *GetTenantUsageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get tenant usage default response
func (o *GetTenantUsageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTenantUsageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
