package calculo

import (
	"testing"

	"github.com/DjalmaFO/calculo_distancia_entre_pontos_golang/pontos"
	"github.com/DjalmaFO/calculo_distancia_entre_pontos_golang/validator"
)

func TestCoordenadosPonto(t *testing.T) {
	dados := []validator.Validator{
		{Valor: pontos.Ponto{Latitude: -47.6577, Longitude: 179.9999}, Esperado: true},
		{Valor: pontos.Ponto{Latitude: -90.6577, Longitude: 179.9999}, Esperado: false},
		{Valor: pontos.Ponto{Latitude: -47.6577, Longitude: 180.0000}, Esperado: true},
		{Valor: pontos.Ponto{Latitude: -47.6577, Longitude: 180.0001}, Esperado: false},
		{Valor: pontos.Ponto{Latitude: 47.6577, Longitude: -179.9999}, Esperado: true},
		{Valor: pontos.Ponto{Latitude: 90.6577, Longitude: 179.9999}, Esperado: false},
		{Valor: pontos.Ponto{Latitude: 47.6577, Longitude: -180.9999}, Esperado: false},
	}

	if !validator.Validate(dados, pontos.ValidarCoordenadasPonto) {
		t.Errorf("Falha no teste de validar coordenadas")
		t.Fail()
	}
}
