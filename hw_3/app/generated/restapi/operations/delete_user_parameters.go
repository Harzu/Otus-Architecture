// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteUserParams creates a new DeleteUserParams object
// no default values defined in spec.
func NewDeleteUserParams() DeleteUserParams {

	return DeleteUserParams{}
}

// DeleteUserParams contains all the bound params for the delete user operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteUser
type DeleteUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Body []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteUserParams() beforehand.
func (o *DeleteUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate inline body array
			o.Body = body
			if err := o.validateBodyBody(route.Formats); err != nil {
				res = append(res, err)
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// validateBodyBody validates an inline body parameter
func (o *DeleteUserParams) validateBodyBody(formats strfmt.Registry) error {
	bodyIC := o.Body

	var bodyIR []string
	for i, bodyIV := range bodyIC {
		bodyI := bodyIV

		if err := validate.MinLength(fmt.Sprintf("%s.%v", "body", i), "body", bodyI, 1); err != nil {
			return err
		}

		bodyIR = append(bodyIR, bodyI)
	}

	o.Body = bodyIR
	return nil
}