package utils

import (
	"time"

	"github.com/google/uuid"
)

func GenerateIDAndUnixMillis() (uuid.UUID, int64, int64) {
	id, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate UUID: " + err.Error())
	}
	currentMillis := time.Now().UnixMilli()
	return id, currentMillis, currentMillis
}
