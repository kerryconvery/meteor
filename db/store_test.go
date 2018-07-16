package db

import "testing"

func TestWrite(t *testing.T) {
	store := New("../test_data/db")

	defer store.Close()

	writeErr := store.UpdatePosition("movieA", 100)

	if writeErr != nil {
		t.Errorf("Expected no write error but got %s", writeErr)
	}

	position, readErr := store.ReadPosition("movieA")

	if readErr != nil {
		t.Errorf("Expected no read error but got %s", readErr)
	}

	if position != 100 {
		t.Errorf("Expected position to be 100 but got %d", position)
	}
}
