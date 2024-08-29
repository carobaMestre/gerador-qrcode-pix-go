package utils

import (
	"fmt"
)

// GenerateEMV gera uma tag EMV a partir do ID e do valor
func GenerateEMV(id, value string) string {
	length := fmt.Sprintf("%02d", len(value))
	return id + length + value
}
