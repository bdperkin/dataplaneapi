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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetSpoeMessagesParams creates a new GetSpoeMessagesParams object
// no default values defined in spec.
func NewGetSpoeMessagesParams() GetSpoeMessagesParams {

	return GetSpoeMessagesParams{}
}

// GetSpoeMessagesParams contains all the bound params for the get spoe messages operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSpoeMessages
type GetSpoeMessagesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Spoe scope
	  Required: true
	  In: query
	*/
	Scope string
	/*Spoe file name
	  Required: true
	  In: query
	*/
	Spoe string
	/*ID of the transaction where we want to add the operation. Cannot be used when version is specified.
	  In: query
	*/
	TransactionID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSpoeMessagesParams() beforehand.
func (o *GetSpoeMessagesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qScope, qhkScope, _ := qs.GetOK("scope")
	if err := o.bindScope(qScope, qhkScope, route.Formats); err != nil {
		res = append(res, err)
	}

	qSpoe, qhkSpoe, _ := qs.GetOK("spoe")
	if err := o.bindSpoe(qSpoe, qhkSpoe, route.Formats); err != nil {
		res = append(res, err)
	}

	qTransactionID, qhkTransactionID, _ := qs.GetOK("transaction_id")
	if err := o.bindTransactionID(qTransactionID, qhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindScope binds and validates parameter Scope from query.
func (o *GetSpoeMessagesParams) bindScope(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("scope", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("scope", "query", raw); err != nil {
		return err
	}

	o.Scope = raw

	return nil
}

// bindSpoe binds and validates parameter Spoe from query.
func (o *GetSpoeMessagesParams) bindSpoe(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("spoe", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("spoe", "query", raw); err != nil {
		return err
	}

	o.Spoe = raw

	return nil
}

// bindTransactionID binds and validates parameter TransactionID from query.
func (o *GetSpoeMessagesParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TransactionID = &raw

	return nil
}
