package db

import (
	"testing"
)

func TestUpdatePosition(t *testing.T) {
	store := NewMediaStore("../test_data/db")

	defer store.Close()

	writeErr := store.UpdatePosition("movieA", 100)

	if writeErr != nil {
		t.Errorf("Expected no write error but got %s", writeErr)
	}

	data, readErr := store.Read("movieA")

	if readErr != nil {
		t.Errorf("Expected no read error but got %s", readErr)
	}

	if data.Position != 100 {
		t.Errorf("Expected position to be 100 but got %d", data.Position)
	}
}

func TestUpdateDuration(t *testing.T) {
	store := NewMediaStore("../test_data/db")

	defer store.Close()

	writeErr := store.UpdateDuration("movieA", 10)

	if writeErr != nil {
		t.Errorf("Expected no write error but got %s", writeErr)
	}

	data, readErr := store.Read("movieA")

	if readErr != nil {
		t.Errorf("Expected no read error but got %s", readErr)
	}

	if data.Duration != 10 {
		t.Errorf("Expected position to be 100 but got %d", data.Duration)
	}
}
