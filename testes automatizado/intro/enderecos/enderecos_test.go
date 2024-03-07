package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// precisa começar com Test obrigatoriamente

	enderecoParaTeste := "Rua Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoEndereco(enderecoParaTeste)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diference do esperado, esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}

}
