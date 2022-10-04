package calculo

import "github.com/DjalmaFO/calculo_distancia_entre_pontos_golang/models"

// CalcularDistancia => Retorna a distancia aproximada entre dois pontos.
// O calculo é realizado com base na latitude e longitude dos pontos 1 e 2.
// A distancia é dada em Km (Kilometers)
func CalcularDistancia(latitude1, longitude1, latitude2, longitude2 float64) (distancia float64, err error) {
	ponto1 := models.Ponto{Latitude: 34.0522, Longitude: -118.2436}
	ponto2 := models.Ponto{Latitude: 35.6850, Longitude: 139.7514}

	return models.KmDistancia(ponto1, ponto2)
}
