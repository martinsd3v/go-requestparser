# go-requestparser/parser

Golang package for working with arrays on http requests

###### FRONT

```html
<form method="POST" enctype="multipart/form-data">
    <input type="text" name="User[Name]" value="Jonh">
    <input type="text" name="User[Email]" value="Jonh@mail.com">
    <input type="text" name="User[Pass]" value="123">
    <input type="text" name="User[Contact][0][Phone]" value="123456">
    <input type="text" name="User[Contact][0][Name]" value="Mom">
    <input type="text" name="User[Contact][1][Phone]" value="654321">
    <input type="text" name="User[Contact][1][Name]" value="Daddy">
</form>
```

###### BACK-END
```go
import "github.com/martinsd3v/go-requestparser/parser"

func HandleRequest(rw http.ResponseWriter, request *http.Request) {
	//... HandleRequest
	var RequestDTO struct {
		User struct {
			Name string
			Email string
			Pass int
			Contact []struct{
				Phone string
				Name string
			} 
		}
	}

	parsed := parser.Parser(request, RequestDTO)
	//HandleRequest ...
}
```

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
If you are unhappy to use long `parser`, you can do something like this:
```go
import (
  ps "github.com/martinsd3v/go-requestparser/parser"
)
```

#### My motivations:

This package came out of a personal need, as soon as I started to develop with GoLang I found it difficult to handle arrays received via HTTP requests. Because Go takes a different approach, especially for developers used to Node.js or PHP. This does not happen with JSON requests since data in json format is easily converted to language structure. So to receive multidimensional structure in Go, you would have to work with json, which in itself is not a problem. However if you take into account that the service must be able to function independently of the way in which the data is sent (Json, Query Params, form-data, etc.) this ends up becoming a big problem. Then using this package you can easily convert the received data to a Go structure.

So, if you are looking for a way to submit forms without having to convert your data to json and have easy access to the information on your backend, this package can be useful.

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
Se você não gostar de usar o `parser`, pode fazer algo assim:
```go
import (
  ps "github.com/martinsd3v/go-requestparser/parser"
)
```

#### Minhas motivações:

Esse pacote surgiu de uma necessidade pessoal, assim que comecei a desenvolver com GoLang  me deparei com a dificuldade tratar arrays recebidos via requisições HTTP. Pois Go tem uma abordagem diferente, principalmente para desenvolvedores acostumados com Node.js ou PHP. Isso não acontece com requisições do tipo JSON pois dados no formato json são facilmente convertidas para estruturas da linguagem. Então para receber estrutura multidimensional em Go teria obrigatoriamente que trabalhar com json,  oque por si só não e um problema. Entretanto se levar em consideração que o serviço deve ser capas de funcionar independente da forma em que os dados são enviados ( Json, Query Params, form-data, etc. ) isso acaba se tornando um grande problema. Então utilizando esse pacote você consegue converter facilmente os dados recebidos para uma estrutura em Go.

Portanto, se você está procurando uma maneira de enviar formulários sem ter que converter seus dados para json e ter fácil acesso às informações em seu back-end, este pacote pode ser util.