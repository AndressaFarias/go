package main

import (
	"fmt"
)
	
type ContaCorrente struct{
	titular string
	numAgencia int
	numConta int
	saldo float64
}

func main() {

	contaFulano := ContaCorrente{titular: "Fulano", numAgencia: 123, numConta: 123456, saldo: 123.45}
	fmt.Println(contaFulano)

	contaBeltrano := ContaCorrente{"Beltrano", 456, 456789, 456.89}
	fmt.Println(contaBeltrano)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)

	contaDaCris.titular = "Cris"
	fmt.Println(*contaDaCris)
}