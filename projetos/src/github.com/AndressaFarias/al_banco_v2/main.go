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
	contaDaCris.saldo = 500.0

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)

	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500.0

	fmt.Println(contaDaCris)
	fmt.Println(contaDaCris2)

	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)




}