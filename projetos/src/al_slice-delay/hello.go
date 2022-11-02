package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 5
const delay = 3


func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		op := leOpcao()

		switch op {
		case 1:
			iniciarMonitoramento()
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

}

func exibeIntroducao() {
	nome := "Malulu Cá"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leOpcao() int {
	var opLida int
	fmt.Scan(&opLida)
	fmt.Println("O valor da variável comando é:", opLida)

	return opLida
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://www.alura.com.br/", "https://random-status-code.herokuapp.com/", "https://www.caelum.com.br"}

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
			fmt.Println(" ")
		}
		time.Sleep(delay * time.Second)
		fmt.Println("\n--")
	}
	fmt.Println("****")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problema. Status Code:", resp.StatusCode)
		fmt.Println("Site:", site, "está com problema. Status Code:", resp.Status)
	}
}
