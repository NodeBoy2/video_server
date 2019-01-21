package utils

import (
	uuid "github.com/satori/go.uuid"
)

func NewUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
