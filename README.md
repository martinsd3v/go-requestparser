# go-requestparser/parser

Golang package for working with arrays on http requests form-data

#### Installation
Make sure that Go is installed on your computer.
Type the following command in your terminal:
```bash
go get github.com/martinsd3v/go-requestparser/parser
```

#### Import package in your project
Add following line in your `*.go` file:
```go
import "github.com/martinsd3v/go-requestparser/parser"
```
If you are unhappy to use long `govalidator`, you can do something like this:
```go
import (
  valid "github.com/martinsd3v/go-requestparser/parser"
)
```

#### My motivations:

This package came out of a personal need, as soon as I started the development in golang I soon came across a difficulty in sending data in array format via form-data. Because golang works in a very different way from other languages like javascript and php. to work with arrays it would be mandatory to work with json, that did not please me!

So if you are looking for a way to send forms without having to convert your data to json and have easy access to the information on your back-end, this is the right package.

###### PT-BR

Pacote em golang para trabalhar com arrays em requisições http form-data

#### Instalação
Certifique-se de que Go está instalado em seu computador.
Digite o seguinte comando em seu terminal:

```bash
go get github.com/martinsd3v/go-requestparser/parser
```

#### Importar pacote em seu projeto
Adicione a seguinte linha em seu arquivo `*.go`:
```go
import "github.com/martinsd3v/go-requestparser/parser"
```
Se você não gosta de usar o `go-requestparser/parser` longo, pode fazer algo assim:
```go
import (
  valid "github.com/martinsd3v/go-requestparser/parser"
)
```

#### Minhas motivações:

Esse pacote surgiu de uma necessidade pessoal, assim que começei o desenvolvimento em golang logo me deparei com uma dificuldade em enviar dados em formato array via form-data. Pois golang trabalha de uma forma muito diferente de outras linguagens como javascript e php. para trabalhar com arrays seria necessário obrigatoriamente trabalhar com json, isso não me agradou! 

Portanto se você busca uma forma de enviar formularios sem precisar converter seus dados em json e ter acesso facilmente as informações em seu back-end, esté é o pacote certo.