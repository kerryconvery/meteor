package controllers

import (
	"bytes"
)

// JSONResponse represents a json response
type JSONResponse struct {
	StatusCode  int
	ContentType string
	Body        interface{}
}

// JSONResponse represents a json response
type BinaryResponse struct {
	StatusCode  int
	ContentType string
	Body        *bytes.Buffer
}

// Controller represents a controller
type Controller struct {
}

// JSONResponse returns a json response struct
func (c Controller) JSONResponse(statusCode int, body interface{}) JSONResponse {
	return JSONResponse{StatusCode: statusCode, ContentType: "application/json", Body: body}
}

func (c Controller) BinaryResponse(contentType string, body *bytes.Buffer) BinaryResponse {
	return BinaryResponse{StatusCode: 200, ContentType: contentType, Body: body}
}
