# Go

* Linguagem modularizada
* 25 palavras chaves

# Instalação do Go

Para instalarmos o Go em nosso computador primeiramente precisamos fazer o Download do binário da linguagem adequado.
*  Acesse a página de download oficial da linguagem e faça o download adequado para o seu sistema - https://go.dev/dl/


## Instalando para o Ubuntu

* Faça o download do arquivo .tar do Go;
* Acesse via terminal a pasta em que o arquivo foi baixado;
* Execute o comando a seguir para extraí-lo em sua pasta `/usr/local`:
`sudo tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz`

Precisamos adicionar o caminho `/usr/local/go/bin` no PATH do sistema.
Você pode fazer isto adicionando uma linha extra no arquivo `/etc/profile`, 

* Abra o arquivo /etc/profile com o vim: `sudo vim /etc/profile`

E em seguida adicione a seguinte linha no final do arquivo: `export PATH=$PATH:/usr/local/go/bin`

* Agora execute o comando : `source /etc/profile`

E o Go já deve estar funcionando com sucesso! 
Faça o teste em seu terminal com o comando `go version` e veja se a versão e os comandos do Go foram exibidos corretamente.


# Criando o seu workspace

Para utilizarmos a linguagem de programação Go devemos seguir algumas convenções, e a principal delas é a sua organização de pastas no seu Workspace.

O Workspace padrão do Go é um diretório chamado `/go` que fica na pasta do seu usuário em seu sistema operacional. 
No Windows esta pasta normalmente se encontra em C:/Users/seu-usuario/;
nos sistemas Unix ( MacOS e distribuições do Linux) normalmente se encontra em /home/seu-usuario/.

Dentro deste diretório devem conter as seguintes pastas:

~~~shell
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
~~~

* A bin , onde ficará os compilados do nosso código Go.

* A pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando ao longo de nossos códigos.

* A src, onde escreveremos o código fonte de cada aplicação.

Crie estas pastas como indicado acima, para que o seu Go Workspace fique funcionando corretamente.

# Instalando o Visual Studio Code

## Plugin do Go
Para sua vida seja facilitada quando você está trabalhando com Go no Visual Studio Code, você deve instalar os plugins recomendados por ele. Se você criar qualquer arquivo com a extensão .go, uma pequena pop up surgirá perguntando se você deseja instalar a extensão.


Ao clicar em Show Recommendations você verá no Menu lateral a extensão chamada "Go", do autor lukehoban. Selecione-a clicando em Install e em seguida ao final da instalação em Reload para recarregamos a janela.

No momento em que a janela for recarregada, a extensão será habilitada, mas ela tem algumas dependências, conforme é exibida na notificação do Visual Studio Code. Então, basta clicar em Install All, para instalar todas as dependências, que facilitarão a nossa codificação em Go.


