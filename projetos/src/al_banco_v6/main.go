package main

import (
	"fmt"
	c "al_banco_v6/contas"
)
	


func main() {
	contaDaSilvia := c.ContaCorrente{Titular:"Silvia", Saldo: 300}
	contaDoGustavo := c.ContaCorrente{Titular:"Gustavo", Saldo: 100}


    fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	contaDaSilvia.Transferir(150.,&contaDoGustavo)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}