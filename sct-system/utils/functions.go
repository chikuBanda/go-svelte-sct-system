package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomTransferStatus() string {
	stringSlice := []string{
		"INITIATED",
		"PENDING",
		"COMPLETED",
		"ABORTED",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(stringSlice))
	pick := stringSlice[randomIndex]
	return string(pick)
}
