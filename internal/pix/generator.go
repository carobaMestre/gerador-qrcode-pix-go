package pix

import (
    "fmt"
    "gerador-qrcode-pix-go/internal/utils"
)

// QrCodePixParams define os parâmetros necessários para o QR Code PIX
type QrCodePixParams struct {
    Version       string
    Key           string
    City          string
    Name          string
    Value         float64
    TransactionId string
}

// GeneratePayload gera o payload do QR Code PIX com base nos parâmetros
func GeneratePayload(params QrCodePixParams) string {
    payload := utils.GenerateEMV("00", params.Version)
    payload += utils.GenerateEMV("26", params.Key)
    payload += utils.GenerateEMV("52", "0000")
    payload += utils.GenerateEMV("53", "986")
    payload += utils.GenerateEMV("54", fmt.Sprintf("%.2f", params.Value))
    payload += utils.GenerateEMV("58", "BR")
    payload += utils.GenerateEMV("59", params.Name)
    payload += utils.GenerateEMV("60", params.City)
    payload += utils.GenerateEMV("62", params.TransactionId)

    return payload + utils.GenerateCRC16(payload)
}
