package main

import "fmt"

func boilingPointWaterKelvinToCelsius() int {
	K := 373
	var waterEbulicaoC = K - 273
	return waterEbulicaoC
}

func main() {
	fmt.Printf("O valor de ebulução da água convertida de kelvin para celsius é : %d.", boilingPointWaterKelvinToCelsius())
}
