package models

import (
	"fmt"
	"log"
	m "math"

	"github.com/DjalmaFO/calculo_distancia_entre_pontos_golang/funcs"
)

type Pontos struct {
	PontoOrigem  Ponto
	PontoDestino Ponto
}

type Ponto struct {
	Latitude  float64
	Longitude float64
}

func (ps *Pontos) CalcularDistancia() (distancia float64) {
	// Referências:
	// "https://www.ehow.com.br/calcular-distancia-entre-pontos-latitude-longitude-como_71372/"
	// "https://pt.stackoverflow.com/questions/159969/calcular-dist%C3%A2ncia-entre-dois-pontos-pela-latitude-e-longitude"
	const raioTerra = 6371

	// Converter medida em graus para radianos
	lat1 := ps.PontoOrigem.Latitude * m.Pi / 180 // m = package math
	long1 := ps.PontoOrigem.Longitude * m.Pi / 180
	lat2 := ps.PontoDestino.Latitude * m.Pi / 180
	long2 := ps.PontoDestino.Longitude * m.Pi / 180

	// Diferença entre latitudes
	deltaLat := lat2 - lat1

	// Diferença entre longitudes
	deltaLong := long2 - long1

	// m = package math
	// val1 = [sen²(Δlat/2) + cos(lat1)] x cos(lat2) x sen²(Δlong/2)
	val1 := m.Sin(deltaLat/2)*m.Sin(deltaLat/2) + m.Cos(lat1)*m.Cos(lat2)*(m.Sin(deltaLong/2)*m.Sin(deltaLong/2))

	//val2 = 2 x cot(√val1/√(1−val1)) => "onde "cot" é o inverso da função tangente, indicada como 'tan^−1' em algumas calculadoras."
	val2 := 2 * funcs.Atan2(m.Sqrt(val1), m.Sqrt(1-val1))

	// retorno = R x val2, onde "R" representa o raio da Terra (6,371 km)
	return raioTerra * val2
}

func KmDistancia(p1, p2 Ponto) (kmDistancia float64, err error) {

	if ok := ValidarCoordenadasPonto(p1); !ok {
		return kmDistancia, fmt.Errorf("Cordenadas inválidas para ponto 1: %.4f e %.4f", p1.Latitude, p1.Longitude)
	}

	if ok := ValidarCoordenadasPonto(p2); !ok {
		return kmDistancia, fmt.Errorf("Cordenadas inválidas para ponto 2: %.4f e %.4f", p2.Latitude, p2.Longitude)
	}
	pts := Pontos{p1, p2}

	return pts.CalcularDistancia(), nil
}

func ValidarCoordenadasPonto(value interface{}) (ok bool) {
	var p Ponto
	switch v := value.(type) {
	case Ponto:
		p = value.(Ponto)
	default:
		log.Printf("Função aceita apenas estrutura de Ponto{} e não %v \n", v)
		return false
	}

	if p.Latitude > 90 || p.Latitude < -90 {
		return
	}

	if p.Longitude > 180 || p.Longitude < -180 {
		return
	}

	return true
}
