package tests

import "testing"

func ExpectError(err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func ExpectNoError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Expected no errors but got %s", err)
	}
}
