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

package spoe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// GetSpoeMessageOKCode is the HTTP code returned for type GetSpoeMessageOK
const GetSpoeMessageOKCode int = 200

/*GetSpoeMessageOK Successful operation

swagger:response getSpoeMessageOK
*/
type GetSpoeMessageOK struct {
	/*Spoe configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetSpoeMessageOKBody `json:"body,omitempty"`
}

// NewGetSpoeMessageOK creates GetSpoeMessageOK with default headers values
func NewGetSpoeMessageOK() *GetSpoeMessageOK {

	return &GetSpoeMessageOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe message o k response
func (o *GetSpoeMessageOK) WithConfigurationVersion(configurationVersion int64) *GetSpoeMessageOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe message o k response
func (o *GetSpoeMessageOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe message o k response
func (o *GetSpoeMessageOK) WithPayload(payload *GetSpoeMessageOKBody) *GetSpoeMessageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe message o k response
func (o *GetSpoeMessageOK) SetPayload(payload *GetSpoeMessageOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeMessageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSpoeMessageNotFoundCode is the HTTP code returned for type GetSpoeMessageNotFound
const GetSpoeMessageNotFoundCode int = 404

/*GetSpoeMessageNotFound The specified resource was not found

swagger:response getSpoeMessageNotFound
*/
type GetSpoeMessageNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeMessageNotFound creates GetSpoeMessageNotFound with default headers values
func NewGetSpoeMessageNotFound() *GetSpoeMessageNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetSpoeMessageNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe message not found response
func (o *GetSpoeMessageNotFound) WithConfigurationVersion(configurationVersion int64) *GetSpoeMessageNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe message not found response
func (o *GetSpoeMessageNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe message not found response
func (o *GetSpoeMessageNotFound) WithPayload(payload *models.Error) *GetSpoeMessageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe message not found response
func (o *GetSpoeMessageNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeMessageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetSpoeMessageDefault General Error

swagger:response getSpoeMessageDefault
*/
type GetSpoeMessageDefault struct {
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

// NewGetSpoeMessageDefault creates GetSpoeMessageDefault with default headers values
func NewGetSpoeMessageDefault(code int) *GetSpoeMessageDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetSpoeMessageDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get spoe message default response
func (o *GetSpoeMessageDefault) WithStatusCode(code int) *GetSpoeMessageDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe message default response
func (o *GetSpoeMessageDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe message default response
func (o *GetSpoeMessageDefault) WithConfigurationVersion(configurationVersion int64) *GetSpoeMessageDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe message default response
func (o *GetSpoeMessageDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe message default response
func (o *GetSpoeMessageDefault) WithPayload(payload *models.Error) *GetSpoeMessageDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe message default response
func (o *GetSpoeMessageDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeMessageDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
