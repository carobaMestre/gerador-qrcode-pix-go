# Gerador de QR Code PIX em Go

![Go Version](https://img.shields.io/badge/go-1.23-blue)
![License](https://img.shields.io/github/license/carobaMestre/gerador-qrcode-pix-go)

## Descrição

O **Gerador de QR Code PIX em Go** é uma ferramenta simples e eficiente para criar QR Codes de pagamento via PIX. Este projeto foi desenvolvido com o objetivo de facilitar a geração de QR Codes para transações financeiras, utilizando a linguagem Go.

## Funcionalidades

- **Geração de Payload EMV:** Gera o payload EMV necessário para a criação de QR Codes PIX.
- **Validação de Parâmetros:** Valida os parâmetros fornecidos para garantir a conformidade com as especificações do PIX.
- **Geração de QR Code:** Converte o payload EMV gerado em um QR Code no formato Base64.
- **Checksum CRC16:** Calcula o checksum CRC16 para garantir a integridade dos dados.

## Estrutura do Projeto

```plaintext
gerador-qrcode-pix-go/
├── cmd/
│   └── main.go            # Ponto de entrada do aplicativo
├── internal/
│   ├── pix/
│   │   ├── generator.go    # Lógica para gerar o payload do QR Code PIX
│   │   ├── qrcode.go       # Geração do QR Code em Base64
│   │   └── validator.go    # Validação dos parâmetros do QR Code
│   └── utils/
│       ├── crc16.go        # Funções relacionadas ao cálculo do CRC16
│       └── emv.go          # Funções utilitárias para geração de tags EMV
├── test/
│   └── pix_test.go         # Testes unitários para o pacote pix
├── go.mod                  # Gerenciamento de dependências Go
├── go.sum                  # Hashes das dependências do projeto
└── README.md               # Documentação do projeto
```

## Como Usar

### Requisitos

- [Go 1.23+](https://golang.org/dl/)

### Passos para Instalação

1. Clone o repositório:

   ```sh
   git clone https://github.com/carobaMestre/gerador-qrcode-pix-go.git
   ```

2. Navegue até o diretório do projeto:

   ```sh
   cd gerador-qrcode-pix-go
   ```

3. Instale as dependências:

   ```sh
   go mod tidy
   ```

### Exemplo de Uso

No arquivo `cmd/main.go`, você pode configurar os parâmetros do QR Code PIX e gerar o QR Code:

```go
package main

import (
    "fmt"
    "gerador-qrcode-pix-go/internal/pix"
)

func main() {
    params := pix.QrCodePixParams{
        Version:       "01",
        Key:           "12345678900",
        City:          "Sao Paulo",
        Name:          "Fulano de Tal",
        Value:         100.00,
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
```

### Executar o Projeto

Execute o projeto usando o comando:

```sh
go run cmd/main.go
```

## Contribuindo

Contribuições são bem-vindas! Se você deseja contribuir com este projeto, siga estas etapas:

1. Faça um fork do repositório.
2. Crie uma branch para a sua feature (`git checkout -b feature/AmazingFeature`).
3. Commit suas alterações (`git commit -m 'Add some AmazingFeature'`).
4. Faça o push para a branch (`git push origin feature/AmazingFeature`).
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

## Contato

Desenvolvido por [Illumi Caroba](https://github.com/carobaMestre) - entre em contato!

---