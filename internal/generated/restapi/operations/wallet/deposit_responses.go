// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/oke11o/walletsuro/internal/generated/models"
)

// DepositOKCode is the HTTP code returned for type DepositOK
const DepositOKCode int = 200

/*DepositOK Ok

swagger:response depositOK
*/
type DepositOK struct {

	/*
	  In: Body
	*/
	Payload *models.Wallet `json:"body,omitempty"`
}

// NewDepositOK creates DepositOK with default headers values
func NewDepositOK() *DepositOK {

	return &DepositOK{}
}

// WithPayload adds the payload to the deposit o k response
func (o *DepositOK) WithPayload(payload *models.Wallet) *DepositOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deposit o k response
func (o *DepositOK) SetPayload(payload *models.Wallet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DepositOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DepositInternalServerErrorCode is the HTTP code returned for type DepositInternalServerError
const DepositInternalServerErrorCode int = 500

/*DepositInternalServerError Invalid input

swagger:response depositInternalServerError
*/
type DepositInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.SimpleResponse `json:"body,omitempty"`
}

// NewDepositInternalServerError creates DepositInternalServerError with default headers values
func NewDepositInternalServerError() *DepositInternalServerError {

	return &DepositInternalServerError{}
}

// WithPayload adds the payload to the deposit internal server error response
func (o *DepositInternalServerError) WithPayload(payload *models.SimpleResponse) *DepositInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deposit internal server error response
func (o *DepositInternalServerError) SetPayload(payload *models.SimpleResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DepositInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
