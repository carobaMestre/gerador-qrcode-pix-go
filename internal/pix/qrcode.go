package pix

import (
    "github.com/skip2/go-qrcode"
)

// GenerateQrCode gera o QR Code em formato base64 a partir do payload
func GenerateQrCode(payload string) (string, error) {
    png, err := qrcode.Encode(payload, qrcode.Medium, 256)
    if err != nil {
        return "", err
    }
    return string(png), nil
}
