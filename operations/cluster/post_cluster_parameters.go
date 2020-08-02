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

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/haproxytech/models/v2"
)

// NewPostClusterParams creates a new PostClusterParams object
// no default values defined in spec.
func NewPostClusterParams() PostClusterParams {

	return PostClusterParams{}
}

// PostClusterParams contains all the bound params for the post cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters postCluster
type PostClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*In case of moving to single mode do we keep or clean configuration
	  In: query
	*/
	Configuration *string
	/*
	  Required: true
	  In: body
	*/
	Data *models.ClusterSettings
	/*Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.
	  In: query
	*/
	Version *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostClusterParams() beforehand.
func (o *PostClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qConfiguration, qhkConfiguration, _ := qs.GetOK("configuration")
	if err := o.bindConfiguration(qConfiguration, qhkConfiguration, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ClusterSettings
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("data", "body", ""))
			} else {
				res = append(res, errors.NewParseError("data", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Data = &body
			}
		}
	} else {
		res = append(res, errors.Required("data", "body", ""))
	}
	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindConfiguration binds and validates parameter Configuration from query.
func (o *PostClusterParams) bindConfiguration(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Configuration = &raw

	if err := o.validateConfiguration(formats); err != nil {
		return err
	}

	return nil
}

// validateConfiguration carries on validations for parameter Configuration
func (o *PostClusterParams) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Enum("configuration", "query", *o.Configuration, []interface{}{"keep"}); err != nil {
		return err
	}

	return nil
}

// bindVersion binds and validates parameter Version from query.
func (o *PostClusterParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("version", "query", "int64", raw)
	}
	o.Version = &value

	return nil
}
