// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oke11o/walletsuro/internal/generated/models"
)

// TransferOKCode is the HTTP code returned for type TransferOK
const TransferOKCode int = 200

/*TransferOK Ok

swagger:response transferOK
*/
type TransferOK struct {

	/*
	  In: Body
	*/
	Payload *models.Wallet `json:"body,omitempty"`
}

// NewTransferOK creates TransferOK with default headers values
func NewTransferOK() *TransferOK {

	return &TransferOK{}
}

// WithPayload adds the payload to the transfer o k response
func (o *TransferOK) WithPayload(payload *models.Wallet) *TransferOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transfer o k response
func (o *TransferOK) SetPayload(payload *models.Wallet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransferOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransferBadRequestCode is the HTTP code returned for type TransferBadRequest
const TransferBadRequestCode int = 400

/*TransferBadRequest Bad request

swagger:response transferBadRequest
*/
type TransferBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewTransferBadRequest creates TransferBadRequest with default headers values
func NewTransferBadRequest() *TransferBadRequest {

	return &TransferBadRequest{}
}

// WithPayload adds the payload to the transfer bad request response
func (o *TransferBadRequest) WithPayload(payload *models.SimpleResponse) *TransferBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transfer bad request response
func (o *TransferBadRequest) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransferBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransferForbiddenCode is the HTTP code returned for type TransferForbidden
const TransferForbiddenCode int = 403

/*TransferForbidden Permission denied

swagger:response transferForbidden
*/
type TransferForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewTransferForbidden creates TransferForbidden with default headers values
func NewTransferForbidden() *TransferForbidden {

	return &TransferForbidden{}
}

// WithPayload adds the payload to the transfer forbidden response
func (o *TransferForbidden) WithPayload(payload *models.SimpleResponse) *TransferForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transfer forbidden response
func (o *TransferForbidden) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransferForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransferInternalServerErrorCode is the HTTP code returned for type TransferInternalServerError
const TransferInternalServerErrorCode int = 500

/*TransferInternalServerError Invalid input

swagger:response transferInternalServerError
*/
type TransferInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewTransferInternalServerError creates TransferInternalServerError with default headers values
func NewTransferInternalServerError() *TransferInternalServerError {

	return &TransferInternalServerError{}
}

// WithPayload adds the payload to the transfer internal server error response
func (o *TransferInternalServerError) WithPayload(payload *models.SimpleResponse) *TransferInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transfer internal server error response
func (o *TransferInternalServerError) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransferInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
