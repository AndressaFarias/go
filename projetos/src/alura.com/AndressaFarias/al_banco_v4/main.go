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

// é preciso indicar que a operação será feita no objeto ContaCorrenta que está executando a fnção
// para essa indicação é preciso declarar c *ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo 

	if podeSacar{
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar (valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0{
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
} 

func main() {
	contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia)

	// podemos exibir apenas o saldo
	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Sacar(100))

	fmt.Println(contaDaSilvia.saldo)

	fmt.Println(contaDaSilvia.Depositar(2000))
}