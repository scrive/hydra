// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
	"github.com/ory/hydra/internal/httpclient/models"
)

// PerformOAuth2AuthorizationFlowReader is a Reader for the PerformOAuth2AuthorizationFlow structure.
type PerformOAuth2AuthorizationFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformOAuth2AuthorizationFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewPerformOAuth2AuthorizationFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformOAuth2AuthorizationFlowDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformOAuth2AuthorizationFlowFound creates a PerformOAuth2AuthorizationFlowFound with default headers values
func NewPerformOAuth2AuthorizationFlowFound() *PerformOAuth2AuthorizationFlowFound {
	return &PerformOAuth2AuthorizationFlowFound{}
}

/* PerformOAuth2AuthorizationFlowFound describes a response with status code 302, with default header values.

 Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type PerformOAuth2AuthorizationFlowFound struct {
}

func (o *PerformOAuth2AuthorizationFlowFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth][%d] performOAuth2AuthorizationFlowFound ", 302)
}

func (o *PerformOAuth2AuthorizationFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPerformOAuth2AuthorizationFlowDefault creates a PerformOAuth2AuthorizationFlowDefault with default headers values
func NewPerformOAuth2AuthorizationFlowDefault(code int) *PerformOAuth2AuthorizationFlowDefault {
	return &PerformOAuth2AuthorizationFlowDefault{
		_statusCode: code,
	}
}

/* PerformOAuth2AuthorizationFlowDefault describes a response with status code -1, with default header values.

oAuth2ApiError
*/
type PerformOAuth2AuthorizationFlowDefault struct {
	_statusCode int

	Payload *models.OAuth2APIError
}

// Code gets the status code for the perform o auth2 authorization flow default response
func (o *PerformOAuth2AuthorizationFlowDefault) Code() int {
	return o._statusCode
}

func (o *PerformOAuth2AuthorizationFlowDefault) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth][%d] performOAuth2AuthorizationFlow default  %+v", o._statusCode, o.Payload)
}
func (o *PerformOAuth2AuthorizationFlowDefault) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *PerformOAuth2AuthorizationFlowDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
