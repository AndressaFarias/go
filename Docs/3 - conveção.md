* o programa principal deve se indicado por `package main` no topo do arquivo. Isso indica que a execução do programa iniciária por esse arquivo

* o programa principal deve conter a função principal: 

~~~go
func main(){
    ...
}
~~~

* Um função externa sempre inicia com a primeira letra em maiusculo. Ex `fmt.Println("Olá Mundo")`

* No fim da instrução o ponto e virgula é opcional - recomenda-se que não seja utiizado

* variaveis em go:
    * o tipo é informado depois do nome  da variavel;
        * EX: `var nome string`
    * quando uma variavel não possui valor atrbuido, ela inicia com valor zerado
    * se uma variavel é declarada e não usada a complação gera um erro;
    * é capaz de inferir otipo da variavel, ou seja, nem sempre é preciso declarar o tipo da variavel (string / int / boolean)
        * EX: `var nome string = "malu"`> `var nome = "malu"`
              `var idade int = 6`> `var idade = 6`
    * é possivel descartar a declaração de variavel `var`através da utilização do operador de declaração curta de variaveis `:=`, desta forma, informando ao Go que uma variável está sendo declarada e, ao mesmo tempo, um valor está sendo atribuído a ela.
        * EX: `nome := "malu"`
              `idade := 6`

    - a declaração de variáveis do Go começa com a palavra chave `var`, seguida do nome da variável.
    
    - Não há necessidade de declarar o tipo, embora possa ser necessário.

* concatenação  é feita pela virgula 

*  para capturar o que o usuário escrever no teclado, é possivel usar duas funções do pacote fmt
    * Scanf, do pacote fmt : Recebe como argumento um modificador e um ponteiro para a variável que guardará a entrada do usuário
        EX:
        ~~~go
            var idade int
            fmt.Scanf("%d", &idade)
        ~~~
        
        A variável idade é do tipo inteiro, logo, só pode receber números inteiros. Então, como o Go sabe que a variável só pode receber números inteiros, o modificador `%d`, que representa um número inteiro, deixa de ser necessário. 
        Por isso há uma outra função, que não recebe como argumento o modificador.
    * Scan :  também do pacote fmt. Ela consegue inferir o tipo digitado pelo usuário, baseado no tipo da variável recebida.


* `if` deve sempre retornar um valor booleano

    ~~~go

     if comando == 1 {

    } else if comando == 2 {

    } else if comando == 0 {

    }else{

    }
    ~~~

* as instruções `if` podem ou não ter sua condicional entre parênteses. Normalmente elas são escritas sem os parênteses mesmo

* switch : A palavra reservada switch permite incluir várias alternativas de execução para o programa. Cada alternativa deve utilizar case, e a alternativa executada quando nenhuma condição for atendida é default.

    * No Go no uso do break não é obrigatório. Porém, ele não é obrigatório. Apenas uma alternativa é executada por avaliação switch.

* Exit / Sair  do programa
    ~~~go
        import "os"
        os.Exit(0)
    ~~~ 
     A função Exit fica no pacote `os`, que deve ser importado.


* "net/http" é pacote mais específico à nossa necessidade. Já que nele temos funções para realizar requisições Get e Post.
    Para uma comunicação web com um determinado site é necessário a utilização do pacote "net/http". Que pela própria definição dele tem como objetivo fornecer a implementações de cliente e servidor HTTP


* funções com multiplos retornos
    * Em funções com múltiplos valores, temos que informar os tipos de cada retorno entre parênteses. E declarar variáveis separadas por vírgula para receber os valores de saída. No exemplo, a função retornou valores para três tipos diferentes.


* Para que um loop seja criado emm go é possivel utilizar o `for { }` isso indica que o loop será executado infinitamente

* Array em go possui tamanho fixo
    * declaração 
        ~~~go
            var sites [4]string
        ~~~
    * Quando os arrays são criados, eles assumem os valores padrão para os tipos de seus elementos. No caso, o tipo do array frutas é string e o valor padrão para cada posição do array será vazio. Portanto, o valor impresso será uma string vazia.

    * len(<array|slice>) : Isso aí! A função que usamos para descobrir o tamanho de uma slice é len().

* Slice : Isso aí! Quando é necessário colocar mais elementos do que sua capacidade atual, o slice dobra a capacidade.

* Operador de iteração range
    * O range nos dá acesso a cada item da coleção, nos retornando a posição do item iterado e o próprio item daquela posição

* Trabalhando com o pacote time é possivel obter tempo para realizar um sleep entre as chamadas

* Arquivos em Go
    * Isso aí! Podemos utilizar a função os.Open quando queremos abrir um arquivo em Go.
    
ESCREVER O ARQUIVO


Abrir arquivos
     os.OpenFile - )flags, permissão_
     A função os.OpenFile é mais poderosa que a função os.Open e nos permite configurar determinadas flags para configurar como o arquivo será manipulado.
        alguma dAs flags da função os.OpenFile 
            os.O_CREATE`se o arquivo não existetirele é criado|
            os.O_RDWR : para poder ler e escrever no arquivi
            O_APPEND`: para oder escrever na última linha do arquivo

Escrever em arquivos
# Conversão de tipos com Go

## Conhecendo o pacote strconv
A linguagem Go nos fornece um meio de conversão entre diferentes tipos compatíveis, por exemplo strings para inteiro. Essas conversões são feitas pelo pacote strconv.

## Convertendo um tipo String para um tipo int
https://www.alura.com.br/artigos/conversao-de-tipos-com-go 

# Go: crie uma aplicação web




Trabalhar com tempo e formatá-lo
Exibir o conteúdo de um arquivo


# Função com quantidade de parâmetros indeterminados

* Uma função em Go pode ter um, muitos ou nenhum parâmetro;

* porém é possivel que um função Go receba um número indeterminado de parâmetros. Funções deste tipo são conhecidas como _variadic functions_, ou função variádica.

* Para criar uma variadic function, devemos preceder o tipo do argumento com reticências, conforme o exemplo abaixo:

~~~go
func Somando(numeros ...int) int {
    resultadoDaSoma := 0
    for _, numero := range numeros {
        resultadoDaSoma += numero
    }
    return resultadoDaSoma
}
~~~

* Note o uso das reticências na declaração do parâmetro número: `numeros ...int`. Portanto, podemos criar uma função sem parâmetro, com um, dois, três, ou uma quantidade indeterminada de parâmetros com Go.



----

# Criando um pacote

Podemos criar pacotes para que os mesmo contenham as structs e métodos de um determinado objeto.


Para tonar um campo ou função pública ou privada para outras partes da aplicação, alteramos a primeira letra para minúscula ou maiúscula respectivamente.


Quando queremos deixar pública, deixamos a primeira letra maiúscula e privada minúscula.

Visibilidade é o atributo de uma função ou variável a ser visível para diferentes partes do programa

# Passando um valor ou uma cópia
Métodos são definidos de maneira parecida com funções, mas de uma maneira diferente. 
Existe um (p *Pessoa) que se refere a um ponteiro para a instância criada da estrutura, conforme o exemplo abaixo:

~~~go
package main

import (
    "fmt"
)

type Pessoa struct {
    nome, sobrenome string
}

func (p *Pessoa) ExibirNomeCompleto() string {
    nomeCompleto := p.nome + " " + p.sobrenome
    return nomeCompleto
}

func main() {
    p1 := Pessoa{"Guilherme", "Lima"}
    fmt.Println(p1.ExibirNomeCompleto())
}
~~~


Ao executar este código, temos a saída esperada: `Guilherme Lima`

Nesse caso, passamos para o método o valor encontrado neste ponteiro através do (p *Pessoa).

## Passando uma cópia
Também é possível passar um valor removendo a assinatura do ponteiro (p *Pessoa) para (p Pessoa).

Nesse caso, uma cópia do valor de Pessoa é passada para a função, sem alterar o valor do ponteiro. Portanto, precisamos ficar atentos, já que qualquer alteração que você faça em p se passar por valor não será refletida na fonte p.

Observe este exemplo:

~~~go
package main

import (
    "fmt"
)

type Pessoa struct {
    nome, sobrenome string
}

func (p Pessoa) ExibirNomeCompleto() string {
    p.sobrenome = "Silva"
    nomeCompleto := p.nome + " " + p.sobrenome
    return nomeCompleto
}

func main() {
    p1 := Pessoa{"Guilherme", "Lima"}

    fmt.Println(p1.ExibirNomeCompleto())
    fmt.Println(p1.nome, p1.sobrenome)
}
~~~


Nossa saída será: 
~~~
  Guilherme Silva
  Guilherme Lima
~~~


Observe que alteramos o sobrenome de `p` no método `ExibirNomeCompleto`, mas não foi alterado o valor armazenado no ponteiro. Sendo assim, quando não precisamos alterar o conteúdo de um ponteiro, podemos passar apenas uma cópia

#  Tipos polimórficos em Go

## Interface
Uma interface é a definição de um conjunto de métodos comuns a um ou mais tipos. É o que permite a criação de tipos polimórficos em Go.

Podemos reaproveitar métodos ou funções entre tipos utilizando interface.

# Golang: trabalhando com datas

No Golang temos o pacote `Time`, que contém algumas funções para trabalhar com data e hora.

Quando queremos formatar em Golang, utilizamos a palavra format, seguido de algo bem diferente de outras linguagens de programação. Isso porque os formatos de data e hora parecem com a data e hora reais e não Y-m-d H:i:s 

O que queremos editar	código
dia (com zero)	02
dia (sem o zero)	2
dia da semana (inteiro)	Monday
dia da semana abreviado	Mon
mês com número (com zero)	01
mês com número (sem zero)	1
mês (nome inteiro)	January
mês (nome abreviado)	Jan
ano (inteiro)	2006
ano (abreviado)	06
hora (com zero)	03
hora (sem zero)	3
hora (formato 24 horas)	15
minutos (com zero)	04
minutos (sem zero)	4
segundos (com zero)	05
segundos (sem zero)	5


