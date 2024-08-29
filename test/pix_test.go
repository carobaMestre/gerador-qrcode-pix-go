package test

import (
    "pix-qrcode-generator/internal/pix"
    "testing"
)

func TestGeneratePayload(t *testing.T) {
    params := pix.QrCodePixParams{
        Version:       "01",
        Key:           "12345678900",
        City:          "Sao Paulo",
        Name:          "Fulano de Tal",
        Value:         100.00,
        TransactionId: "abc123",
    }

    expected := "000201..."
    result := pix.GeneratePayload(params)

    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}

func TestValidateParams(t *testing.T) {
    params := pix.QrCodePixParams{
        Version:       "01",
        Key:           "",
        City:          "Sao Paulo",
        Name:          "Fulano de Tal",
        Value:         100.00,
        TransactionId: "abc123",
    }

    err := pix.ValidateParams(params)
    if err == nil {
        t.Errorf("Expected error but got none")
    }
}
