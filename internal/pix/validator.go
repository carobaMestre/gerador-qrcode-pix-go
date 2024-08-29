package pix

import "errors"

// ValidateParams valida os parâmetros de entrada para o QR Code PIX
func ValidateParams(params QrCodePixParams) error {
    if params.Version == "" || params.Key == "" || params.City == "" || params.Name == "" {
        return errors.New("parâmetros obrigatórios estão faltando")
    }
    if params.Value <= 0 {
        return errors.New("o valor deve ser maior que zero")
    }
    return nil
}
