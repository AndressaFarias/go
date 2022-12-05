// Uma função em Go pde tter, um, muitos ou nenhum parâmetro;

package main;

import (
	"fmt"
)

func semParametro() string {
	return "Exemplo de função sem parâmetro!"
}

func umParametro (texto string) string {
	return texto
}

func doisParametros (texto string, numero int) (string, int) {
	return texto, numero
}

// Criando uma função variádica
/*Para criar uma variadic function, devemos preceder o tipo do argumento com reticências, conforme o exemplo abaixo*/

func Somando (numeros ...int) int {
	resultadoDaSoma := 0

	for _, numero := range numeros{
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}


func main(){
	fmt.Println(semParametro())
	fmt.Println(umParametro("Exemplo de função com um parãmetro"))
	fmt.Println(doisParametros("Passando dois parâmetros: essa string e o número", 10))

	fmt.Println("TESTE da funçao SOMANDO")

	fmt.Println(Somando(1))
    fmt.Println(Somando(1,1))
    fmt.Println(Somando(1,1,1))
    fmt.Println(Somando(1,1,2,4))

}