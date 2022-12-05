package main

import (
	"fmt"
)

func main (){
	
fmt.Println("Olá mundo Go!")

numero1 := 1
fmt.Println(numero1)

numero2 := 2
numero3 := numero1 + numero2
fmt.Println(numero3)

var mensagem string
if numero1 < numero2 {
   mensagem = "O número 1 é menor que o número 2"
} else if(numero1 == numero2) {
   mensagem = "Os números 1 e 2 são iguais"
} else {
   mensagem = "O número 1 é maior que o número 2"
}
fmt.Println(mensagem)

for i := 0; i <= 50; i ++ {
  if i % 5 == 0 {
   fmt.Println(i, " é um número divisível por 5")
  }
 }
}