package enderecos

import "strings"

// verifica se o endereço tem um tipo válido e retorna
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoMinuscula, " ")[0]

	enderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoValido = true

		}
	}

	if enderecoValido {
		return strings.Title(primeiraPalavra)
	}
	return "Tipo Inválido"
}
