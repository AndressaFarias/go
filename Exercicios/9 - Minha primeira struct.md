# V1
- Criar o novo projeto chamado banco
- criar um arquivo chamado `main.go`
- informar a qual pacote pertence esse código, em nosos caso ele pertence ao `package main`
- criar a função main nesse arquivo
- criar a _struct_ para conter osa dsos da conta, que são
    - titular
    - numero da agencia
    - numero da conta
    - saldo
- Para criar uma struct, usamos o prefixo `type`, seguido do `nome`da struct, e o sufixo `struct`e declaramos os campos entre chaves;
- detalhe interessante dentro da _struct_não é necessário fazer a declaração de `var` antes de declarar a variavel
- Exiba os valores da struct : `fmt.Println(ContaCorrente{})`
- Instancie duas contas correntes, uma deve declarar atributo e valor e a outra deve declarar apenas o valor
- 

# V2
- Vamos declarar uma nova struct utilizando usando ponteiros
    - Vamos declarar uma variável chamada `contaDaCris`. Ela será do tipo ContaCorrente.
    - Na linha seguinte iremos criar a nova conta: `contaDaCris = new()`
    - Dentro dos parênteses passaremos o tipo, que é `ContaCorrente`.
- A partir de então, conseguiremos atribuir os valores acrescentando um ponto (.) após `contaDaCris` e o nome do campo na sequência. Então atribuiremos o valor desse campo. `contaDaCris.titular = "Cris"`
- Execute o código
- Porém, aparecerá um erro,

~~~go
# command-line-arguments
./main.go:23:16: cannot use new(ContaCorrente) (value of type *ContaCorrente) as type ContaCorrente in assignment
~~~
Não podemos utilizar nosso tipo `ContaCorrente` já atribuindo um valor. 
Isso porque temos uma variável do tipo `ContaCorrente`, a `contaDaCris`. Mas o código não entendeu que o tipo da variável é o mesmo que está sendo passado para `new()`.
Precisamos identificar que se tratam do mesmo tipo.
Para conseguir dizer que `contaDaCris` aponta para uma `ContaCorrente`, colocaremos um asterisco na frente. 
Teremos `var contaDaCris *ContaCorrente`.
- Imprima o conteúdo da struct
- Veremos no terminal: `&{Cris 0 0 0}`, o quê difere dos retornos das outras instanciações. 
Mas por que o `"&"` e por que temos que apontar para `ContaCorrente` com asterisco? Nas outras alternativas de uso da struct, o código entendia que a conta corrente do cliente apontava para um tipo, uma estrutura de conta. 
A `contaDaCris` precisa apontar para `ContaCorrente`. O código em Go não entenderá corretamente se tirarmos e asterisco e salvarmos a aplicação, pois ficará sem entender para onde `new` está apontando, se é a conta corrente da Cris ou uma nova conta.
O último detalhe será o `"&"` que aparece no terminal quando fazemos a impressão do que temos. Isso porque seguindo o mesmo raciocínio do nosso exemplo, `&{Cris 0 0 500}` indica que os campos são um conteúdo que está dentro. Mas "&" não é interessante para nós, somente o conteúdo.

Por isso, colocaremos o asterisco em contaDaCris também na hora da impressão.

