package tests

import "testing"

// ExpectError is used when you are expecting an error
func ExpectError(err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

//ExpectNoError is used when you don't expect an error
func ExpectNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected no errors but got %s", err)
	}
}

// ExpectedContentType is used to check that the current content type was received
func ExpectedContentType(expected, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected content type %s but got %s", expected, actual)
	}
}

// ExpectedStatusCode is used to check that the current status code was received
func ExpectedStatusCode(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected status code %d but got %d", expected, actual)
	}
}
