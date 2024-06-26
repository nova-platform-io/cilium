// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package statedb

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetStatedbDumpReader is a Reader for the GetStatedbDump structure.
type GetStatedbDumpReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetStatedbDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatedbDumpOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /statedb/dump] GetStatedbDump", response, response.Code())
	}
}

// NewGetStatedbDumpOK creates a GetStatedbDumpOK with default headers values
func NewGetStatedbDumpOK(writer io.Writer) *GetStatedbDumpOK {
	return &GetStatedbDumpOK{

		Payload: writer,
	}
}

/*
GetStatedbDumpOK describes a response with status code 200, with default header values.

Success
*/
type GetStatedbDumpOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this get statedb dump o k response has a 2xx status code
func (o *GetStatedbDumpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get statedb dump o k response has a 3xx status code
func (o *GetStatedbDumpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get statedb dump o k response has a 4xx status code
func (o *GetStatedbDumpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get statedb dump o k response has a 5xx status code
func (o *GetStatedbDumpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get statedb dump o k response a status code equal to that given
func (o *GetStatedbDumpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get statedb dump o k response
func (o *GetStatedbDumpOK) Code() int {
	return 200
}

func (o *GetStatedbDumpOK) Error() string {
	return fmt.Sprintf("[GET /statedb/dump][%d] getStatedbDumpOK  %+v", 200, o.Payload)
}

func (o *GetStatedbDumpOK) String() string {
	return fmt.Sprintf("[GET /statedb/dump][%d] getStatedbDumpOK  %+v", 200, o.Payload)
}

func (o *GetStatedbDumpOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetStatedbDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
