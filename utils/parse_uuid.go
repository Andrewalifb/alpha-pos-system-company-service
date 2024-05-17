package utils

import (
	"github.com/google/uuid"
)

func ParseUUID(s string) *uuid.UUID {
	if s == "" {
		return nil
	}
	id, err := uuid.Parse(s)
	if err != nil {
		panic(err)
	}
	return &id
}
