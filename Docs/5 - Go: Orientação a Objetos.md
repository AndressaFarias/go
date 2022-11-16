* variaveis guardam as informações na memória do computador.

# Struct

Estrutura - é um tipo que vai definir as caracteristicas de um objeto;
Para criar uma struct, usamos o prefixo `type`, seguido do `nome`da struct, e o sufixo `struct`e declaramos os campos entre chaves, confomre o exemplo abaixo:

~~~go
type ContaCorrente struct{
    titular string
    numAgencia int
    numConta int
    saldo float64
}
~~~

* Para instanciar a `struct`é preciso declarar `NomeStruct{}`

* Utilizando Struct
    para instanciar a struct com valores devemos delcarar e atribuir o valor

    ~~~go
    contaFulano := ContaCorrente{titular: "Fulano", numAgencia: 123, numConta: 123456, saldo: 123,45}
    ~~~

    outro modo é passar apenas os valores a serem atribuidos, porém esse modo só pode ser usado quadno todos os atributos da struct forem ser preenchidos
     ~~~go
    contaBeltrano := ContaCorrente{"Fulano", 456, numConta: 987564, saldo: 741,25}
    ~~~


----

Para referenciarmos o ponteiro no momento da criação do tipo, podemos colocar entre parênteses logo após func e antes de Sacar() a inscrição (c *ContaCorrente), o que significa que quando a função for chamada, o código apontará para a conta que chama. Nesse caso, quando chamarmos a função, não precisaremos especificar que nos tratamos da conta de um cliente ou outro.


