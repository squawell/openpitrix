// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// PostV1AppsNoContentCode is the HTTP code returned for type PostV1AppsNoContent
const PostV1AppsNoContentCode int = 204

/*PostV1AppsNoContent App succesfully created.

swagger:response postV1AppsNoContent
*/
type PostV1AppsNoContent struct {
}

// NewPostV1AppsNoContent creates PostV1AppsNoContent with default headers values
func NewPostV1AppsNoContent() *PostV1AppsNoContent {
	return &PostV1AppsNoContent{}
}

// WriteResponse to the client
func (o *PostV1AppsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PostV1AppsBadRequestCode is the HTTP code returned for type PostV1AppsBadRequest
const PostV1AppsBadRequestCode int = 400

/*PostV1AppsBadRequest App couldn't have been created.

swagger:response postV1AppsBadRequest
*/
type PostV1AppsBadRequest struct {
}

// NewPostV1AppsBadRequest creates PostV1AppsBadRequest with default headers values
func NewPostV1AppsBadRequest() *PostV1AppsBadRequest {
	return &PostV1AppsBadRequest{}
}

// WriteResponse to the client
func (o *PostV1AppsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// PostV1AppsInternalServerErrorCode is the HTTP code returned for type PostV1AppsInternalServerError
const PostV1AppsInternalServerErrorCode int = 500

/*PostV1AppsInternalServerError An unexpected error occured.

swagger:response postV1AppsInternalServerError
*/
type PostV1AppsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostV1AppsInternalServerError creates PostV1AppsInternalServerError with default headers values
func NewPostV1AppsInternalServerError() *PostV1AppsInternalServerError {
	return &PostV1AppsInternalServerError{}
}

// WithPayload adds the payload to the post v1 apps internal server error response
func (o *PostV1AppsInternalServerError) WithPayload(payload *models.Error) *PostV1AppsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 apps internal server error response
func (o *PostV1AppsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1AppsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
