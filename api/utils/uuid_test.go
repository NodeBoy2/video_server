package utils

import "testing"

func TestNewUUID(t *testing.T) {
	_, err := NewUUID()
	if err != nil {
		t.Errorf("Create UUID Error: %v", err)
	}
}
