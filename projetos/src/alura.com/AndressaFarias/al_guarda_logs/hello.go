package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 2
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
			fmt.Println("Exibindo Logs ...")
			imprimeLogs()
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
	nome := "== Malu == "
	versao := 1.8

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Print("\n")

}

func leOpcao() int {
	var opLida int
	fmt.Scan(&opLida)
	fmt.Println("O valor da variável comando é:", opLida)

	return opLida
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()
	
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
	// passamos a capturar osegundo retorno
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site,true)
	} else {
		fmt.Println("Site:", site, "está com problema. Status Code:", resp.StatusCode)
		registraLog(site,false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for{
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites,linha)

		if err == io.EOF{
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool){
	
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)

	if err != nil{
		fmt.Println("Ocorreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

// ai ler o nosso arquivo de logs e imprimir no nosso terminal:
func imprimeLogs(){

	fmt.Println("entrando na função imprimeLog")
	
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(arquivo))
}