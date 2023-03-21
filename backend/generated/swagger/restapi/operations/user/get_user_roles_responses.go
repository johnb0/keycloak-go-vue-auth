// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"test_iam/generated/swagger/models"
)

// GetUserRolesOKCode is the HTTP code returned for type GetUserRolesOK
const GetUserRolesOKCode int = 200

/*
GetUserRolesOK Success

swagger:response getUserRolesOK
*/
type GetUserRolesOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewGetUserRolesOK creates GetUserRolesOK with default headers values
func NewGetUserRolesOK() *GetUserRolesOK {

	return &GetUserRolesOK{}
}

// WithPayload adds the payload to the get user roles o k response
func (o *GetUserRolesOK) WithPayload(payload []string) *GetUserRolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user roles o k response
func (o *GetUserRolesOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
GetUserRolesDefault Unexpected error.

swagger:response getUserRolesDefault
*/
type GetUserRolesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserRolesDefault creates GetUserRolesDefault with default headers values
func NewGetUserRolesDefault(code int) *GetUserRolesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUserRolesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get user roles default response
func (o *GetUserRolesDefault) WithStatusCode(code int) *GetUserRolesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get user roles default response
func (o *GetUserRolesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get user roles default response
func (o *GetUserRolesDefault) WithPayload(payload *models.Error) *GetUserRolesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user roles default response
func (o *GetUserRolesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRolesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}