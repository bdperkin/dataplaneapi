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

// CreateSpoeGroupCreatedCode is the HTTP code returned for type CreateSpoeGroupCreated
const CreateSpoeGroupCreatedCode int = 201

/*CreateSpoeGroupCreated Spoe groups created

swagger:response createSpoeGroupCreated
*/
type CreateSpoeGroupCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SpoeGroup `json:"body,omitempty"`
}

// NewCreateSpoeGroupCreated creates CreateSpoeGroupCreated with default headers values
func NewCreateSpoeGroupCreated() *CreateSpoeGroupCreated {

	return &CreateSpoeGroupCreated{}
}

// WithPayload adds the payload to the create spoe group created response
func (o *CreateSpoeGroupCreated) WithPayload(payload *models.SpoeGroup) *CreateSpoeGroupCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe group created response
func (o *CreateSpoeGroupCreated) SetPayload(payload *models.SpoeGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeGroupCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSpoeGroupBadRequestCode is the HTTP code returned for type CreateSpoeGroupBadRequest
const CreateSpoeGroupBadRequestCode int = 400

/*CreateSpoeGroupBadRequest Bad request

swagger:response createSpoeGroupBadRequest
*/
type CreateSpoeGroupBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSpoeGroupBadRequest creates CreateSpoeGroupBadRequest with default headers values
func NewCreateSpoeGroupBadRequest() *CreateSpoeGroupBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateSpoeGroupBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create spoe group bad request response
func (o *CreateSpoeGroupBadRequest) WithConfigurationVersion(configurationVersion int64) *CreateSpoeGroupBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create spoe group bad request response
func (o *CreateSpoeGroupBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create spoe group bad request response
func (o *CreateSpoeGroupBadRequest) WithPayload(payload *models.Error) *CreateSpoeGroupBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe group bad request response
func (o *CreateSpoeGroupBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeGroupBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateSpoeGroupConflictCode is the HTTP code returned for type CreateSpoeGroupConflict
const CreateSpoeGroupConflictCode int = 409

/*CreateSpoeGroupConflict The specified resource already exists

swagger:response createSpoeGroupConflict
*/
type CreateSpoeGroupConflict struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateSpoeGroupConflict creates CreateSpoeGroupConflict with default headers values
func NewCreateSpoeGroupConflict() *CreateSpoeGroupConflict {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateSpoeGroupConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create spoe group conflict response
func (o *CreateSpoeGroupConflict) WithConfigurationVersion(configurationVersion int64) *CreateSpoeGroupConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create spoe group conflict response
func (o *CreateSpoeGroupConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create spoe group conflict response
func (o *CreateSpoeGroupConflict) WithPayload(payload *models.Error) *CreateSpoeGroupConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe group conflict response
func (o *CreateSpoeGroupConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeGroupConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Configuration-Version

	configurationVersion := swag.FormatInt64(o.ConfigurationVersion)
	if configurationVersion != "" {
		rw.Header().Set("Configuration-Version", configurationVersion)
	}

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateSpoeGroupDefault General Error

swagger:response createSpoeGroupDefault
*/
type CreateSpoeGroupDefault struct {
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

// NewCreateSpoeGroupDefault creates CreateSpoeGroupDefault with default headers values
func NewCreateSpoeGroupDefault(code int) *CreateSpoeGroupDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateSpoeGroupDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the create spoe group default response
func (o *CreateSpoeGroupDefault) WithStatusCode(code int) *CreateSpoeGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create spoe group default response
func (o *CreateSpoeGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create spoe group default response
func (o *CreateSpoeGroupDefault) WithConfigurationVersion(configurationVersion int64) *CreateSpoeGroupDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create spoe group default response
func (o *CreateSpoeGroupDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create spoe group default response
func (o *CreateSpoeGroupDefault) WithPayload(payload *models.Error) *CreateSpoeGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create spoe group default response
func (o *CreateSpoeGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateSpoeGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
