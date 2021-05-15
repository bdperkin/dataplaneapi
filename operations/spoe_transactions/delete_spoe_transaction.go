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

package spoe_transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSpoeTransactionHandlerFunc turns a function with the right signature into a delete spoe transaction handler
type DeleteSpoeTransactionHandlerFunc func(DeleteSpoeTransactionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteSpoeTransactionHandlerFunc) Handle(params DeleteSpoeTransactionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteSpoeTransactionHandler interface for that can handle valid delete spoe transaction params
type DeleteSpoeTransactionHandler interface {
	Handle(DeleteSpoeTransactionParams, interface{}) middleware.Responder
}

// NewDeleteSpoeTransaction creates a new http.Handler for the delete spoe transaction operation
func NewDeleteSpoeTransaction(ctx *middleware.Context, handler DeleteSpoeTransactionHandler) *DeleteSpoeTransaction {
	return &DeleteSpoeTransaction{Context: ctx, Handler: handler}
}

/*DeleteSpoeTransaction swagger:route DELETE /services/haproxy/spoe_transactions/{id} SpoeTransactions deleteSpoeTransaction

Delete a transaction

Deletes a transaction.

*/
type DeleteSpoeTransaction struct {
	Context *middleware.Context
	Handler DeleteSpoeTransactionHandler
}

func (o *DeleteSpoeTransaction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteSpoeTransactionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
