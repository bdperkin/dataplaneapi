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

// GetSpoeAgentOKCode is the HTTP code returned for type GetSpoeAgentOK
const GetSpoeAgentOKCode int = 200

/*GetSpoeAgentOK Successful operation

swagger:response getSpoeAgentOK
*/
type GetSpoeAgentOK struct {
	/*Spoe configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetSpoeAgentOKBody `json:"body,omitempty"`
}

// NewGetSpoeAgentOK creates GetSpoeAgentOK with default headers values
func NewGetSpoeAgentOK() *GetSpoeAgentOK {

	return &GetSpoeAgentOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe agent o k response
func (o *GetSpoeAgentOK) WithConfigurationVersion(configurationVersion int64) *GetSpoeAgentOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe agent o k response
func (o *GetSpoeAgentOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe agent o k response
func (o *GetSpoeAgentOK) WithPayload(payload *GetSpoeAgentOKBody) *GetSpoeAgentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe agent o k response
func (o *GetSpoeAgentOK) SetPayload(payload *GetSpoeAgentOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeAgentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetSpoeAgentNotFoundCode is the HTTP code returned for type GetSpoeAgentNotFound
const GetSpoeAgentNotFoundCode int = 404

/*GetSpoeAgentNotFound The specified resource was not found

swagger:response getSpoeAgentNotFound
*/
type GetSpoeAgentNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSpoeAgentNotFound creates GetSpoeAgentNotFound with default headers values
func NewGetSpoeAgentNotFound() *GetSpoeAgentNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetSpoeAgentNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get spoe agent not found response
func (o *GetSpoeAgentNotFound) WithConfigurationVersion(configurationVersion int64) *GetSpoeAgentNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe agent not found response
func (o *GetSpoeAgentNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe agent not found response
func (o *GetSpoeAgentNotFound) WithPayload(payload *models.Error) *GetSpoeAgentNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe agent not found response
func (o *GetSpoeAgentNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeAgentNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetSpoeAgentDefault General Error

swagger:response getSpoeAgentDefault
*/
type GetSpoeAgentDefault struct {
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

// NewGetSpoeAgentDefault creates GetSpoeAgentDefault with default headers values
func NewGetSpoeAgentDefault(code int) *GetSpoeAgentDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetSpoeAgentDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get spoe agent default response
func (o *GetSpoeAgentDefault) WithStatusCode(code int) *GetSpoeAgentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get spoe agent default response
func (o *GetSpoeAgentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get spoe agent default response
func (o *GetSpoeAgentDefault) WithConfigurationVersion(configurationVersion int64) *GetSpoeAgentDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get spoe agent default response
func (o *GetSpoeAgentDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get spoe agent default response
func (o *GetSpoeAgentDefault) WithPayload(payload *models.Error) *GetSpoeAgentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get spoe agent default response
func (o *GetSpoeAgentDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSpoeAgentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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