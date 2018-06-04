package controllers

// JSONResponse represents a json response
type JSONResponse struct {
	StatusCode  int
	ContentType string
	Body        interface{}
}

// Controller represents a controller
type Controller struct {
}

// JSONResponse returns a json response struct
func (c Controller) JSONResponse(statusCode int, body interface{}) JSONResponse {
	return JSONResponse{StatusCode: statusCode, ContentType: "application/json", Body: body}
}
