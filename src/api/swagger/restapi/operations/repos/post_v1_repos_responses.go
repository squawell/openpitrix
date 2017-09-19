// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/src/api/swagger/models"
)

// PostV1ReposNoContentCode is the HTTP code returned for type PostV1ReposNoContent
const PostV1ReposNoContentCode int = 204

/*PostV1ReposNoContent Repo succesfully created.

swagger:response postV1ReposNoContent
*/
type PostV1ReposNoContent struct {
}

// NewPostV1ReposNoContent creates PostV1ReposNoContent with default headers values
func NewPostV1ReposNoContent() *PostV1ReposNoContent {
	return &PostV1ReposNoContent{}
}

// WriteResponse to the client
func (o *PostV1ReposNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PostV1ReposBadRequestCode is the HTTP code returned for type PostV1ReposBadRequest
const PostV1ReposBadRequestCode int = 400

/*PostV1ReposBadRequest Repo couldn't have been created.

swagger:response postV1ReposBadRequest
*/
type PostV1ReposBadRequest struct {
}

// NewPostV1ReposBadRequest creates PostV1ReposBadRequest with default headers values
func NewPostV1ReposBadRequest() *PostV1ReposBadRequest {
	return &PostV1ReposBadRequest{}
}

// WriteResponse to the client
func (o *PostV1ReposBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// PostV1ReposInternalServerErrorCode is the HTTP code returned for type PostV1ReposInternalServerError
const PostV1ReposInternalServerErrorCode int = 500

/*PostV1ReposInternalServerError An unexpected error occured.

swagger:response postV1ReposInternalServerError
*/
type PostV1ReposInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostV1ReposInternalServerError creates PostV1ReposInternalServerError with default headers values
func NewPostV1ReposInternalServerError() *PostV1ReposInternalServerError {
	return &PostV1ReposInternalServerError{}
}

// WithPayload adds the payload to the post v1 repos internal server error response
func (o *PostV1ReposInternalServerError) WithPayload(payload *models.Error) *PostV1ReposInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post v1 repos internal server error response
func (o *PostV1ReposInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostV1ReposInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
