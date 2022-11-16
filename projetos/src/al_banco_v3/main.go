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

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo 

	if podeSacar{
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}else{
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia)

	contaDaSilvia.Sacar(200)

	// podemos exibir apenas o saldo
	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(200))

	fmt.Println(contaDaSilvia.saldo)

}