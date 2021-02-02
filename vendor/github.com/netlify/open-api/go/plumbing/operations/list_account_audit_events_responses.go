// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// ListAccountAuditEventsReader is a Reader for the ListAccountAuditEvents structure.
type ListAccountAuditEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountAuditEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAccountAuditEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAccountAuditEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAccountAuditEventsOK creates a ListAccountAuditEventsOK with default headers values
func NewListAccountAuditEventsOK() *ListAccountAuditEventsOK {
	return &ListAccountAuditEventsOK{}
}

/*ListAccountAuditEventsOK handles this case with default header values.

OK
*/
type ListAccountAuditEventsOK struct {
	Payload []*models.AuditLog
}

func (o *ListAccountAuditEventsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/audit][%d] listAccountAuditEventsOK  %+v", 200, o.Payload)
}

func (o *ListAccountAuditEventsOK) GetPayload() []*models.AuditLog {
	return o.Payload
}

func (o *ListAccountAuditEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountAuditEventsDefault creates a ListAccountAuditEventsDefault with default headers values
func NewListAccountAuditEventsDefault(code int) *ListAccountAuditEventsDefault {
	return &ListAccountAuditEventsDefault{
		_statusCode: code,
	}
}

/*ListAccountAuditEventsDefault handles this case with default header values.

error
*/
type ListAccountAuditEventsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list account audit events default response
func (o *ListAccountAuditEventsDefault) Code() int {
	return o._statusCode
}

func (o *ListAccountAuditEventsDefault) Error() string {
	return fmt.Sprintf("[GET /accounts/{account_id}/audit][%d] listAccountAuditEvents default  %+v", o._statusCode, o.Payload)
}

func (o *ListAccountAuditEventsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAccountAuditEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}