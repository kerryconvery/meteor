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

//TextResponse represents a text response
type TextResponse struct {
	StatusCode  int
	ContentType string
	Body        string
}

// Controller represents a controller
type Controller struct {
}

// JSONResponse returns a json response struct
func (c Controller) JSONResponse(statusCode int, body interface{}) JSONResponse {
	return JSONResponse{StatusCode: statusCode, ContentType: "application/json", Body: body}
}

// BinaryResponse returns a binary response struct
func (c Controller) BinaryResponse(contentType string, body *bytes.Buffer) BinaryResponse {
	return BinaryResponse{StatusCode: 200, ContentType: contentType, Body: body}
}

// TextResponse returns a text response struct
func (c Controller) TextResponse(statusCode int, body string) TextResponse {
	return TextResponse{StatusCode: statusCode, ContentType: "text/html; charset=utf-8", Body: body}
}
