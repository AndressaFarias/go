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


* Para que um loop seja criado emm go é possivel utilizar o `for { }` isso indica qeu o loop será executado infinitamente

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
    