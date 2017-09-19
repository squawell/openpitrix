// Code generated by go-swagger; DO NOT EDIT.

package app_runtimes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// GetV1AppruntimesOKCode is the HTTP code returned for type GetV1AppruntimesOK
const GetV1AppruntimesOKCode int = 200

/*GetV1AppruntimesOK A list of app runtimes

swagger:response getV1AppruntimesOK
*/
type GetV1AppruntimesOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetV1AppruntimesOKBody `json:"body,omitempty"`
}

// NewGetV1AppruntimesOK creates GetV1AppruntimesOK with default headers values
func NewGetV1AppruntimesOK() *GetV1AppruntimesOK {
	return &GetV1AppruntimesOK{}
}

// WithPayload adds the payload to the get v1 appruntimes o k response
func (o *GetV1AppruntimesOK) WithPayload(payload *models.GetV1AppruntimesOKBody) *GetV1AppruntimesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 appruntimes o k response
func (o *GetV1AppruntimesOK) SetPayload(payload *models.GetV1AppruntimesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AppruntimesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1AppruntimesInternalServerErrorCode is the HTTP code returned for type GetV1AppruntimesInternalServerError
const GetV1AppruntimesInternalServerErrorCode int = 500

/*GetV1AppruntimesInternalServerError An unexpected error occured.

swagger:response getV1AppruntimesInternalServerError
*/
type GetV1AppruntimesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetV1AppruntimesInternalServerError creates GetV1AppruntimesInternalServerError with default headers values
func NewGetV1AppruntimesInternalServerError() *GetV1AppruntimesInternalServerError {
	return &GetV1AppruntimesInternalServerError{}
}

// WithPayload adds the payload to the get v1 appruntimes internal server error response
func (o *GetV1AppruntimesInternalServerError) WithPayload(payload *models.Error) *GetV1AppruntimesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 appruntimes internal server error response
func (o *GetV1AppruntimesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1AppruntimesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
