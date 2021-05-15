// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// GetAllStorageMapFilesOKCode is the HTTP code returned for type GetAllStorageMapFilesOK
const GetAllStorageMapFilesOKCode int = 200

/*GetAllStorageMapFilesOK Successful operation

swagger:response getAllStorageMapFilesOK
*/
type GetAllStorageMapFilesOK struct {

	/*
	  In: Body
	*/
	Payload models.Maps `json:"body,omitempty"`
}

// NewGetAllStorageMapFilesOK creates GetAllStorageMapFilesOK with default headers values
func NewGetAllStorageMapFilesOK() *GetAllStorageMapFilesOK {

	return &GetAllStorageMapFilesOK{}
}

// WithPayload adds the payload to the get all storage map files o k response
func (o *GetAllStorageMapFilesOK) WithPayload(payload models.Maps) *GetAllStorageMapFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all storage map files o k response
func (o *GetAllStorageMapFilesOK) SetPayload(payload models.Maps) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllStorageMapFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Maps{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllStorageMapFilesNotFoundCode is the HTTP code returned for type GetAllStorageMapFilesNotFound
const GetAllStorageMapFilesNotFoundCode int = 404

/*GetAllStorageMapFilesNotFound The specified resource was not found

swagger:response getAllStorageMapFilesNotFound
*/
type GetAllStorageMapFilesNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllStorageMapFilesNotFound creates GetAllStorageMapFilesNotFound with default headers values
func NewGetAllStorageMapFilesNotFound() *GetAllStorageMapFilesNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetAllStorageMapFilesNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get all storage map files not found response
func (o *GetAllStorageMapFilesNotFound) WithConfigurationVersion(configurationVersion int64) *GetAllStorageMapFilesNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get all storage map files not found response
func (o *GetAllStorageMapFilesNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get all storage map files not found response
func (o *GetAllStorageMapFilesNotFound) WithPayload(payload *models.Error) *GetAllStorageMapFilesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all storage map files not found response
func (o *GetAllStorageMapFilesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllStorageMapFilesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetAllStorageMapFilesDefault General Error

swagger:response getAllStorageMapFilesDefault
*/
type GetAllStorageMapFilesDefault struct {
	_statusCode int
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAllStorageMapFilesDefault creates GetAllStorageMapFilesDefault with default headers values
func NewGetAllStorageMapFilesDefault(code int) *GetAllStorageMapFilesDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetAllStorageMapFilesDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get all storage map files default response
func (o *GetAllStorageMapFilesDefault) WithStatusCode(code int) *GetAllStorageMapFilesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get all storage map files default response
func (o *GetAllStorageMapFilesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get all storage map files default response
func (o *GetAllStorageMapFilesDefault) WithConfigurationVersion(configurationVersion int64) *GetAllStorageMapFilesDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get all storage map files default response
func (o *GetAllStorageMapFilesDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get all storage map files default response
func (o *GetAllStorageMapFilesDefault) WithPayload(payload *models.Error) *GetAllStorageMapFilesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all storage map files default response
func (o *GetAllStorageMapFilesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllStorageMapFilesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}