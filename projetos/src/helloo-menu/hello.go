package main 

import "fmt"

func main(){
	nome := "Malu"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:",versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
    fmt.Println("0 - Sair do Programa")

	var op int 
	fmt.Scan(&op)

	fmt.Println("O valor da variável comando é:", op)
}