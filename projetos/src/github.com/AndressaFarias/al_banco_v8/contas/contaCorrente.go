package contas

import (
	cl "al_banco_v8/clientes"
)

type ContaCorrente struct {
	Titular              cl.Titular
	NumAgencia, NumConta int
	saldo                float64
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia <= c.saldo && valorDaTransferencia >= 0 {
		c.Sacar(valorDaTransferencia)
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

// é preciso indicar que a operação será feita no objeto ContaCorrenta que está executando a fnção
// para essa indicação é preciso declarar c *ContaCorrente
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
