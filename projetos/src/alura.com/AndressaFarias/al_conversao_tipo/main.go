package main

import (
    "fmt"
	"strconv"
)

func main() {
    salario := "1000"
    aumento := "20"

    fmt.Println("O salário é", salario, "e o aumento de", aumento + "%")

	// Tentei converter um tipo string para um tipo int, usamos o método Atoi do pacote strconv da seguinte forma:

	salarioConvertido, err := strconv.Atoi(salario)
    if err != nil {
        fmt.Println(err)
    }
    aumentoConvertido, err := strconv.Atoi(aumento)
    if err != nil {
        fmt.Println(err)
    }

	fmt.Println("O salário é", salario, "e o aumento de", aumento + "%")

	novoSalario := (((salarioConvertido * aumentoConvertido) / 100) + salarioConvertido)
    fmt.Println("O novo salário é", novoSalario)


	// PARA CONVERTER INT EM STRING
	numero := 2019
	numeroConvertido := strconv.Itoa(numero)
	fmt.Println(numeroConvertido)
}