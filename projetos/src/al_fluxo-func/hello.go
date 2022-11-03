package main 

import "fmt"
import "os"

func main(){

	exibeIntroducao()
	exibeMenu()
	op := leOpcao()

	switch op {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
		os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
		os.Exit(-1)
    }
}

func exibeIntroducao(){
	nome := "Malulu Cá"
	versao := 1.2

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:",versao)
}

func exibeMenu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
    fmt.Println("0 - Sair do Programa")
}

func leOpcao() int {
	var opLida int
	fmt.Scan(&opLida)
	fmt.Println("O valor da variável comando é:", opLida)

	return opLida;
}