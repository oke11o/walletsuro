// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DepositOKCode is the HTTP code returned for type DepositOK
const DepositOKCode int = 200

/*DepositOK Success

swagger:response depositOK
*/
type DepositOK struct {
}

// NewDepositOK creates DepositOK with default headers values
func NewDepositOK() *DepositOK {

	return &DepositOK{}
}

// WriteResponse to the client
func (o *DepositOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DepositMethodNotAllowedCode is the HTTP code returned for type DepositMethodNotAllowed
const DepositMethodNotAllowedCode int = 405

/*DepositMethodNotAllowed Invalid input

swagger:response depositMethodNotAllowed
*/
type DepositMethodNotAllowed struct {
}

// NewDepositMethodNotAllowed creates DepositMethodNotAllowed with default headers values
func NewDepositMethodNotAllowed() *DepositMethodNotAllowed {

	return &DepositMethodNotAllowed{}
}

// WriteResponse to the client
func (o *DepositMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
