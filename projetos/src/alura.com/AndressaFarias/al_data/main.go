package main

import (
	"fmt"
	"time"
)

func main() {

	//modo hardcode de exibir uma data
	data := "10/Jun/2019 - 08:30:45"
	fmt.Println(data)

	//import do pacote time

	data2 := time.Now()
	fmt.Println(data2)
	// Observe que temos informações demais aqui e este não é o resultado que nós gostaríamos.

	data3 := time.Now()
	fmt.Println(data3.Format(("02/Jan/2006 15:04:05 ")))

}
