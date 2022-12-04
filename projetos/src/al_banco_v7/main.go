package main

import (
	"fmt"
	c "al_banco_v7/contas"
	cl "al_banco_v7/clientes"
)
	


func main() {
	// V1 criação de objeto com struct aninhada
	contaDaSilvia := c.ContaCorrente{Titular: cl.Titular{
		Nome : "Silvia",
		CPF : "123.123.123-12",
		Profissao : "Desenvolvedora"},
		NumAgencia: 3,
		NumConta : 22,
		saldo : 100}
	
	// V2 criação de obejto com struct aninhada

	clienteGustavo := cl.Titular{Nome: "Gustavo", CPF: "123.123.123.12", Profissao: "Arquiteto"}
	contaDoGustavo := c.ContaCorrente{clienteGustavo, 1, 99, 100.}


    fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDoGustavo.saldo)
}