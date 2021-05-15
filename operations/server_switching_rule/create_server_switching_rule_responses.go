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

package server_switching_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// CreateServerSwitchingRuleCreatedCode is the HTTP code returned for type CreateServerSwitchingRuleCreated
const CreateServerSwitchingRuleCreatedCode int = 201

/*CreateServerSwitchingRuleCreated Server Switching Rule created

swagger:response createServerSwitchingRuleCreated
*/
type CreateServerSwitchingRuleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ServerSwitchingRule `json:"body,omitempty"`
}

// NewCreateServerSwitchingRuleCreated creates CreateServerSwitchingRuleCreated with default headers values
func NewCreateServerSwitchingRuleCreated() *CreateServerSwitchingRuleCreated {

	return &CreateServerSwitchingRuleCreated{}
}

// WithPayload adds the payload to the create server switching rule created response
func (o *CreateServerSwitchingRuleCreated) WithPayload(payload *models.ServerSwitchingRule) *CreateServerSwitchingRuleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server switching rule created response
func (o *CreateServerSwitchingRuleCreated) SetPayload(payload *models.ServerSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerSwitchingRuleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServerSwitchingRuleAcceptedCode is the HTTP code returned for type CreateServerSwitchingRuleAccepted
const CreateServerSwitchingRuleAcceptedCode int = 202

/*CreateServerSwitchingRuleAccepted Configuration change accepted and reload requested

swagger:response createServerSwitchingRuleAccepted
*/
type CreateServerSwitchingRuleAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.ServerSwitchingRule `json:"body,omitempty"`
}

// NewCreateServerSwitchingRuleAccepted creates CreateServerSwitchingRuleAccepted with default headers values
func NewCreateServerSwitchingRuleAccepted() *CreateServerSwitchingRuleAccepted {

	return &CreateServerSwitchingRuleAccepted{}
}

// WithReloadID adds the reloadId to the create server switching rule accepted response
func (o *CreateServerSwitchingRuleAccepted) WithReloadID(reloadID string) *CreateServerSwitchingRuleAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create server switching rule accepted response
func (o *CreateServerSwitchingRuleAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create server switching rule accepted response
func (o *CreateServerSwitchingRuleAccepted) WithPayload(payload *models.ServerSwitchingRule) *CreateServerSwitchingRuleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server switching rule accepted response
func (o *CreateServerSwitchingRuleAccepted) SetPayload(payload *models.ServerSwitchingRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerSwitchingRuleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateServerSwitchingRuleBadRequestCode is the HTTP code returned for type CreateServerSwitchingRuleBadRequest
const CreateServerSwitchingRuleBadRequestCode int = 400

/*CreateServerSwitchingRuleBadRequest Bad request

swagger:response createServerSwitchingRuleBadRequest
*/
type CreateServerSwitchingRuleBadRequest struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerSwitchingRuleBadRequest creates CreateServerSwitchingRuleBadRequest with default headers values
func NewCreateServerSwitchingRuleBadRequest() *CreateServerSwitchingRuleBadRequest {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateServerSwitchingRuleBadRequest{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create server switching rule bad request response
func (o *CreateServerSwitchingRuleBadRequest) WithConfigurationVersion(configurationVersion int64) *CreateServerSwitchingRuleBadRequest {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server switching rule bad request response
func (o *CreateServerSwitchingRuleBadRequest) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server switching rule bad request response
func (o *CreateServerSwitchingRuleBadRequest) WithPayload(payload *models.Error) *CreateServerSwitchingRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server switching rule bad request response
func (o *CreateServerSwitchingRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerSwitchingRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateServerSwitchingRuleConflictCode is the HTTP code returned for type CreateServerSwitchingRuleConflict
const CreateServerSwitchingRuleConflictCode int = 409

/*CreateServerSwitchingRuleConflict The specified resource already exists

swagger:response createServerSwitchingRuleConflict
*/
type CreateServerSwitchingRuleConflict struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateServerSwitchingRuleConflict creates CreateServerSwitchingRuleConflict with default headers values
func NewCreateServerSwitchingRuleConflict() *CreateServerSwitchingRuleConflict {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateServerSwitchingRuleConflict{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the create server switching rule conflict response
func (o *CreateServerSwitchingRuleConflict) WithConfigurationVersion(configurationVersion int64) *CreateServerSwitchingRuleConflict {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server switching rule conflict response
func (o *CreateServerSwitchingRuleConflict) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server switching rule conflict response
func (o *CreateServerSwitchingRuleConflict) WithPayload(payload *models.Error) *CreateServerSwitchingRuleConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server switching rule conflict response
func (o *CreateServerSwitchingRuleConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerSwitchingRuleConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*CreateServerSwitchingRuleDefault General Error

swagger:response createServerSwitchingRuleDefault
*/
type CreateServerSwitchingRuleDefault struct {
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

// NewCreateServerSwitchingRuleDefault creates CreateServerSwitchingRuleDefault with default headers values
func NewCreateServerSwitchingRuleDefault(code int) *CreateServerSwitchingRuleDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &CreateServerSwitchingRuleDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) WithStatusCode(code int) *CreateServerSwitchingRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) WithConfigurationVersion(configurationVersion int64) *CreateServerSwitchingRuleDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) WithPayload(payload *models.Error) *CreateServerSwitchingRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create server switching rule default response
func (o *CreateServerSwitchingRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServerSwitchingRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
