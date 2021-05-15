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

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	"github.com/haproxytech/client-native/v2/models"
)

// DeleteBackendAcceptedCode is the HTTP code returned for type DeleteBackendAccepted
const DeleteBackendAcceptedCode int = 202

/*DeleteBackendAccepted Configuration change accepted and reload requested

swagger:response deleteBackendAccepted
*/
type DeleteBackendAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`
}

// NewDeleteBackendAccepted creates DeleteBackendAccepted with default headers values
func NewDeleteBackendAccepted() *DeleteBackendAccepted {

	return &DeleteBackendAccepted{}
}

// WithReloadID adds the reloadId to the delete backend accepted response
func (o *DeleteBackendAccepted) WithReloadID(reloadID string) *DeleteBackendAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the delete backend accepted response
func (o *DeleteBackendAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WriteResponse to the client
func (o *DeleteBackendAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(202)
}

// DeleteBackendNoContentCode is the HTTP code returned for type DeleteBackendNoContent
const DeleteBackendNoContentCode int = 204

/*DeleteBackendNoContent Backend deleted

swagger:response deleteBackendNoContent
*/
type DeleteBackendNoContent struct {
}

// NewDeleteBackendNoContent creates DeleteBackendNoContent with default headers values
func NewDeleteBackendNoContent() *DeleteBackendNoContent {

	return &DeleteBackendNoContent{}
}

// WriteResponse to the client
func (o *DeleteBackendNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteBackendNotFoundCode is the HTTP code returned for type DeleteBackendNotFound
const DeleteBackendNotFoundCode int = 404

/*DeleteBackendNotFound The specified resource was not found

swagger:response deleteBackendNotFound
*/
type DeleteBackendNotFound struct {
	/*Configuration file version

	  Default: 0
	*/
	ConfigurationVersion int64 `json:"Configuration-Version"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteBackendNotFound creates DeleteBackendNotFound with default headers values
func NewDeleteBackendNotFound() *DeleteBackendNotFound {

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &DeleteBackendNotFound{

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithConfigurationVersion adds the configurationVersion to the delete backend not found response
func (o *DeleteBackendNotFound) WithConfigurationVersion(configurationVersion int64) *DeleteBackendNotFound {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete backend not found response
func (o *DeleteBackendNotFound) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete backend not found response
func (o *DeleteBackendNotFound) WithPayload(payload *models.Error) *DeleteBackendNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete backend not found response
func (o *DeleteBackendNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBackendNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

/*DeleteBackendDefault General Error

swagger:response deleteBackendDefault
*/
type DeleteBackendDefault struct {
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

// NewDeleteBackendDefault creates DeleteBackendDefault with default headers values
func NewDeleteBackendDefault(code int) *DeleteBackendDefault {
	if code <= 0 {
		code = 500
	}

	var (
		// initialize headers with default values

		configurationVersionDefault = int64(0)
	)

	return &DeleteBackendDefault{
		_statusCode: code,

		ConfigurationVersion: configurationVersionDefault,
	}
}

// WithStatusCode adds the status to the delete backend default response
func (o *DeleteBackendDefault) WithStatusCode(code int) *DeleteBackendDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete backend default response
func (o *DeleteBackendDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithConfigurationVersion adds the configurationVersion to the delete backend default response
func (o *DeleteBackendDefault) WithConfigurationVersion(configurationVersion int64) *DeleteBackendDefault {
	o.ConfigurationVersion = configurationVersion
	return o
}

// SetConfigurationVersion sets the configurationVersion to the delete backend default response
func (o *DeleteBackendDefault) SetConfigurationVersion(configurationVersion int64) {
	o.ConfigurationVersion = configurationVersion
}

// WithPayload adds the payload to the delete backend default response
func (o *DeleteBackendDefault) WithPayload(payload *models.Error) *DeleteBackendDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete backend default response
func (o *DeleteBackendDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBackendDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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
