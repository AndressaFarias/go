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
