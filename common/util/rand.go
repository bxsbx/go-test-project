package util

import (
	"github.com/google/uuid"
)

func GetRandUUID() string {
	random, _ := uuid.NewRandom()
	return random.String()
}
