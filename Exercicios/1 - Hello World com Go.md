* Crie a pasta `hello` -  irá conter o nosso primeiro script em Go. 
    * Crie a pasta/hello dentro da pasta src do Go Workspace

~~~sh
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
~~~ 

📌 Todos os projetos que você for desenvolver na linguagem Go ficarão em suas próprias pastas no diretório src.

* Cre o `hello.go` dentro pasta /hello:

~~~sh
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
           └── hello.go
~~~

E neste arquivo vamos criar o nosso primeiro programa.

* Como todo projeto em Go, precisamos definir qual será o pacote inicial com a instrução:

~~~go
//hello.go
package main
~~~

* Precisamos definir a função principal do programa, a nossa função main:

~~~go
//hello.go
package main

func main(){

}
~~~

* Desejamos exibir uma mensagem, precisamos importar o pacote `fmt`, que contêm as funções de formatação da linguagem inclusive a função que imprime uma mensagem, a função `fmt.Println()`:

~~~go
//hello.go
package main

import "fmt"

func main(){

}
~~~

~~~go
//hello.go
package main

import "fmt"

func main(){
    fmt.Println("Hello World com Go!")
}
~~~ 

Para executar nosso código, basta utilizarmos o comando `go run hello.go` no terminal dentro da pasta que contêm nosso arquivo com o código fonte do programa que o executável será automaticamente criado e executado :

~~~sh
// Terminal
go run hello.go

Hello World com Go!
~~~
