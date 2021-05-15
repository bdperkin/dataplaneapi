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

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetRuntimeMapEntryParams creates a new GetRuntimeMapEntryParams object
// no default values defined in spec.
func NewGetRuntimeMapEntryParams() GetRuntimeMapEntryParams {

	return GetRuntimeMapEntryParams{}
}

// GetRuntimeMapEntryParams contains all the bound params for the get runtime map entry operation
// typically these are obtained from a http.Request
//
// swagger:parameters getRuntimeMapEntry
type GetRuntimeMapEntryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Map id
	  Required: true
	  In: path
	*/
	ID string
	/*Mapfile attribute storage_name
	  Required: true
	  In: query
	*/
	Map string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRuntimeMapEntryParams() beforehand.
func (o *GetRuntimeMapEntryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qMap, qhkMap, _ := qs.GetOK("map")
	if err := o.bindMap(qMap, qhkMap, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetRuntimeMapEntryParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ID = raw

	return nil
}

// bindMap binds and validates parameter Map from query.
func (o *GetRuntimeMapEntryParams) bindMap(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("map", "query", "")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("map", "query", raw); err != nil {
		return err
	}

	o.Map = raw

	return nil
}
