package contas

type ContaCorrente struct{
	Titular string
	NumAgencia int
	NumConta int
	Saldo float64
}


func (c *ContaCorrente) Transferir (valorDaTransferencia float64, contaDestino *ContaCorrente) (bool){
	if (valorDaTransferencia <= c.Saldo && valorDaTransferencia >= 0){
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
	
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo 

	if podeSacar{
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar (valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0{
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso", c.Saldo
	} else {
		return "Valor do depósito menor que zero", c.Saldo
	}
} 