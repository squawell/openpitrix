// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// GetReposOKCode is the HTTP code returned for type GetReposOK
const GetReposOKCode int = 200

/*GetReposOK A list of repos

swagger:response getReposOK
*/
type GetReposOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetReposOKBody `json:"body,omitempty"`
}

// NewGetReposOK creates GetReposOK with default headers values
func NewGetReposOK() *GetReposOK {
	return &GetReposOK{}
}

// WithPayload adds the payload to the get repos o k response
func (o *GetReposOK) WithPayload(payload *models.GetReposOKBody) *GetReposOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos o k response
func (o *GetReposOK) SetPayload(payload *models.GetReposOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReposInternalServerErrorCode is the HTTP code returned for type GetReposInternalServerError
const GetReposInternalServerErrorCode int = 500

/*GetReposInternalServerError An unexpected error occured.

swagger:response getReposInternalServerError
*/
type GetReposInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReposInternalServerError creates GetReposInternalServerError with default headers values
func NewGetReposInternalServerError() *GetReposInternalServerError {
	return &GetReposInternalServerError{}
}

// WithPayload adds the payload to the get repos internal server error response
func (o *GetReposInternalServerError) WithPayload(payload *models.Error) *GetReposInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get repos internal server error response
func (o *GetReposInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReposInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
