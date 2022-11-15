# Variáveis em Go

## Declarando variáveis em GO

Utilizamos a palavra 'var` + o nome da variável que queremos referenciar + seu tipo.

~~~go
package main

import (
    "fmt"
)

var idade int = 15

func main() {
    fmt.Println(idade)
}
~~~

## Escopo de variáveis

Uma variável declarada dentro de uma estrutura de repetição como o **if, for, switch** ou qualquer outro bloco, não poderá ser acessada fora desse escopo, ou seja, só poderá ser acessadas apenas por escopos mais internos.

Já uma variável ou função criada fora de um bloco é atribuída ao escopo do Package e pode ser acessada de qualquer arquivo dentro do mesmo pacote.

~~~go
package main

import (
    "fmt"
)

func main() {
    var idade int = 15
    fmt.Println(idade)
}
~~~

Porém devemos omitir o tipo da variável int(inteiro), pois isso já será inferido.

Mesmo o Go sendo uma linguagem compilada rápida, de tipagem estática, ou seja, o compilador verifica os tipos usados em dados e variáveis para garantir que sempre está sendo usado um tipo que é esperado em todas as situações, não é aconselhável que isso aconteça, Por exemplo nome :="guilherme" e depois nome =10

## Mais exemplos de declaração de variáveis
Podemos criar mais de uma variável ao mesmo tempo, podendo especificar o tipo ou omitindo:

~~~go
package main

import (
    "fmt"
)

func main() {
    var idade, dia int = 15, 29
    fmt.Println(idade, dia)
}
~~~

Podemos criar mais de uma variável também com tipos diferentes:

~~~go
package main

import (
    "fmt"
)

func main() {
    var (
        idade  int
        altura float32
        nome   string
    )

    idade = 15
    altura = 1.78
    nome = "Guilherme"

    fmt.Println(nome, idade, altura)
}
~~~


Ou, deixar o Go inferir os tipos, para ter o mesmo resultado:

~~~go

package main

import (
    "fmt"
)

func main() {
    var (
        idade  = 15
        altura = 1.78
        nome   = "Guilherme"
    )
    fmt.Println(nome, idade, altura)
}
~~~

Posso também utilizar o **Short variable declarations**, ou seja, declarações curtas de variáveis utilizando o sinal : seguido do =.

~~~go
package main

import (
    "fmt"
)

func main() {

    idade := 15
    altura := 1.78
    nome := "Guilherme"
    fmt.Println(nome, idade, altura)
}
~~~

Está última abordagem com := (dois pontos, igual), usamos apenas quando vamos criar e atribuir um valor dentro do escopo de uma função.

Quando a variável receber um novo valor, devemos utilizar apenas o sinal de = (igual). Caso contrário, receberemos a mensagem dizendo que não há novas variáveis seguido do :=

## Não atribuir um valor ao criar uma variável

Quando criamos uma variável e não atribuímos um valor a ela, o _Go_ se encarrega de atribuir um valor.

~~~
----------------------------------------------
| Tipo         | valor atribuído pelo Go     |
|----------    |-------------------------    |
| int          | 0                           |
| bool         | false                       |
| string       | " "                         |
| float        | 0                           |
| struct       | {}                          |
----------------------------------------------
~~~

É sempre bom fazermos nós mesmos a inicialização das variáveis para não correr nenhum risco dos padrões da linguagem mudarem e nosso programa parar de funcionar.

Porém, existe uma maneira de denotar um ponteiro nulo que, essencialmente, não aponta para nenhum lugar, em Go temo o `nil`

_forma errada de declarar_:

~~~go
package main

import (
    "fmt"
)

func main() {
    a := nil
    fmt.Println(a)
}
~~~


> Não vai compilar.
 O compilador imprimirá o seguinte erro: use of untyped nil, que significa uso não tipado do nil.

_forma correta de declarar_

~~~go
package main

import (
    "fmt"
)

func main() {
    var s *string = nil
    fmt.Println(s)
}
~~~

> Neste caso, o programa compila e retorna <nil> como esperado.


Não podemos fazer a comparação de variaveis , iniciadas com zero, que sejam de tipos diferentes, exemplo:

~~~go
package main

import (
    "fmt"
)

func main() {
    var f float64
    var i int 

    fmt.Println(f==i)
}

~~~