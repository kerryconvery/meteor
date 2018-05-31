package controllers

// HTTPResponse represents an http response
type HTTPResponse struct {
	StatusCode  int
	ContentType string
	Error       error
	Body        interface{}
}

// Controller represents a controller
type Controller struct {
}

// JSONResponse returns a json response struct
func (c Controller) JSONResponse(statusCode int, body interface{}) HTTPResponse {
	return c.Response(statusCode, "application/json", body)
}

// ErrorResponse returns an error response struct
func (c Controller) ErrorResponse(statusCode int, err error) HTTPResponse {
	return HTTPResponse{StatusCode: statusCode, Error: err}
}

// Response returns a response struct
func (c Controller) Response(statusCode int, contentType string, body interface{}) HTTPResponse {
	return HTTPResponse{StatusCode: statusCode, ContentType: contentType, Body: body}
}
