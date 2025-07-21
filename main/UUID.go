package main

import (
	"github.com/google/uuid"
)

func getStringUUID(uuid *uuid.UUID) string {
	return uuid.String()
}

func getStringToUUID(UUID string) (uuid.UUID, error) {
	return uuid.Parse(UUID)
}
