package calculo

import "github.com/DjalmaFO/calculo_distancia_entre_pontos_golang/pontos"

// CalcularDistancia => Retorna a distancia aproximada entre dois pontos.
// O calculo é realizado com base na latitude e longitude dos pontos 1 e 2.
// A distancia é dada em Km (Kilometers)
func CalcularDistancia(latitude1, longitude1, latitude2, longitude2 float64) (distancia float64, err error) {
	ponto1 := pontos.Ponto{Latitude: latitude1, Longitude: longitude1}
	ponto2 := pontos.Ponto{Latitude: latitude2, Longitude: longitude2}

	return pontos.KmDistancia(ponto1, ponto2)
}
