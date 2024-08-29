package main

import (
	"fmt"
	"gerador-qrcode-pix-go/internal/pix"
)

func main() {
	params := pix.QrCodePixParams{
		Version:       "01",
		Key:           "12345678900",
		City:          "Montes Claros",
		Name:          "Caroba Mestre",
		Value:         5000.00,
		TransactionId: "abc123",
	}

	if err := pix.ValidateParams(params); err != nil {
		fmt.Println("Erro de validação:", err)
		return
	}

	payload := pix.GeneratePayload(params)
	qrCode, err := pix.GenerateQrCode(payload)
	if err != nil {
		fmt.Println("Erro ao gerar QR Code:", err)
		return
	}

	fmt.Println("Payload do QR Code PIX:", payload)
	fmt.Println("QR Code Base64:", qrCode)
}
