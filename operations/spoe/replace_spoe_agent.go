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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ReplaceSpoeAgentHandlerFunc turns a function with the right signature into a replace spoe agent handler
type ReplaceSpoeAgentHandlerFunc func(ReplaceSpoeAgentParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceSpoeAgentHandlerFunc) Handle(params ReplaceSpoeAgentParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceSpoeAgentHandler interface for that can handle valid replace spoe agent params
type ReplaceSpoeAgentHandler interface {
	Handle(ReplaceSpoeAgentParams, interface{}) middleware.Responder
}

// NewReplaceSpoeAgent creates a new http.Handler for the replace spoe agent operation
func NewReplaceSpoeAgent(ctx *middleware.Context, handler ReplaceSpoeAgentHandler) *ReplaceSpoeAgent {
	return &ReplaceSpoeAgent{Context: ctx, Handler: handler}
}

/*ReplaceSpoeAgent swagger:route PUT /services/haproxy/spoe/spoe_agents/{name} Spoe replaceSpoeAgent

Replace a SPOE agent

Replaces a SPOE agent configuration in one SPOE scope.

*/
type ReplaceSpoeAgent struct {
	Context *middleware.Context
	Handler ReplaceSpoeAgentHandler
}

func (o *ReplaceSpoeAgent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceSpoeAgentParams()

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
