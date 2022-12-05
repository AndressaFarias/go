package main

import (
	c "al_banco_v8/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoGustavo := c.ContaPoupanca{}
	contaDoGustavo.Depositar(100)
	PagarBoleto(&contaDoGustavo, 60)

	fmt.Println(contaDoGustavo.ObterSaldo())

	contaDaSilvia := c.ContaCorrente{}
	contaDaSilvia.Depositar(500)
	PagarBoleto(&contaDaSilvia, 400)
	

	fmt.Println(contaDaSilvia.ObterSaldo())

}
