// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"transbox/ports/http/gen/models"
)

// CreateEntityCreatedCode is the HTTP code returned for type CreateEntityCreated
const CreateEntityCreatedCode int = 201

/*CreateEntityCreated entity created

swagger:response createEntityCreated
*/
type CreateEntityCreated struct {
}

// NewCreateEntityCreated creates CreateEntityCreated with default headers values
func NewCreateEntityCreated() *CreateEntityCreated {

	return &CreateEntityCreated{}
}

// WriteResponse to the client
func (o *CreateEntityCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// CreateEntityBadRequestCode is the HTTP code returned for type CreateEntityBadRequest
const CreateEntityBadRequestCode int = 400

/*CreateEntityBadRequest bad request

swagger:response createEntityBadRequest
*/
type CreateEntityBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateEntityBadRequest creates CreateEntityBadRequest with default headers values
func NewCreateEntityBadRequest() *CreateEntityBadRequest {

	return &CreateEntityBadRequest{}
}

// WithPayload adds the payload to the create entity bad request response
func (o *CreateEntityBadRequest) WithPayload(payload *models.Error) *CreateEntityBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create entity bad request response
func (o *CreateEntityBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEntityBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEntityInternalServerErrorCode is the HTTP code returned for type CreateEntityInternalServerError
const CreateEntityInternalServerErrorCode int = 500

/*CreateEntityInternalServerError internal error

swagger:response createEntityInternalServerError
*/
type CreateEntityInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateEntityInternalServerError creates CreateEntityInternalServerError with default headers values
func NewCreateEntityInternalServerError() *CreateEntityInternalServerError {

	return &CreateEntityInternalServerError{}
}

// WithPayload adds the payload to the create entity internal server error response
func (o *CreateEntityInternalServerError) WithPayload(payload *models.Error) *CreateEntityInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create entity internal server error response
func (o *CreateEntityInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEntityInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}