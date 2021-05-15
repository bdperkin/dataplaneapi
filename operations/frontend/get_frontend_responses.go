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

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// GetFrontendOKCode is the HTTP code returned for type GetFrontendOK
const GetFrontendOKCode int = 200

/*GetFrontendOK Successful operation

swagger:response getFrontendOK
*/
type GetFrontendOK struct {
	/*Configuration file version

	 */
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *GetFrontendOKBody `json:"body,omitempty"`
}

// NewGetFrontendOK creates GetFrontendOK with default headers values
func NewGetFrontendOK() *GetFrontendOK {

	return &GetFrontendOK{}
}

// WithConfigurationVersion adds the configurationVersion to the get frontend o k response
func (o *GetFrontendOK) WithConfigurationVersion(configurationVersion int64) *GetFrontendOK {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get frontend o k response
func (o *GetFrontendOK) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get frontend o k response
func (o *GetFrontendOK) WithPayload(payload *GetFrontendOKBody) *GetFrontendOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend o k response
func (o *GetFrontendOK) SetPayload(payload *GetFrontendOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetFrontendNotFoundCode is the HTTP code returned for type GetFrontendNotFound
const GetFrontendNotFoundCode int = 404

/*GetFrontendNotFound The specified resource was not found

swagger:response getFrontendNotFound
*/
type GetFrontendNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFrontendNotFound creates GetFrontendNotFound with default headers values
func NewGetFrontendNotFound() *GetFrontendNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetFrontendNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the get frontend not found response
func (o *GetFrontendNotFound) WithConfigurationVersion(configurationVersion int64) *GetFrontendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get frontend not found response
func (o *GetFrontendNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get frontend not found response
func (o *GetFrontendNotFound) WithPayload(payload *models.Error) *GetFrontendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend not found response
func (o *GetFrontendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*GetFrontendDefault General Error

swagger:response getFrontendDefault
*/
type GetFrontendDefault struct {
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

// NewGetFrontendDefault creates GetFrontendDefault with default headers values
func NewGetFrontendDefault(code int) *GetFrontendDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &GetFrontendDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the get frontend default response
func (o *GetFrontendDefault) WithStatusCode(code int) *GetFrontendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get frontend default response
func (o *GetFrontendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the get frontend default response
func (o *GetFrontendDefault) WithConfigurationVersion(configurationVersion int64) *GetFrontendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the get frontend default response
func (o *GetFrontendDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the get frontend default response
func (o *GetFrontendDefault) WithPayload(payload *models.Error) *GetFrontendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get frontend default response
func (o *GetFrontendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFrontendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
