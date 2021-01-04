package main

import (
	"github.com/gin-gonic/gin"
	"github.com/martinsd3v/go-requestparser/parser"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		//01 - Create a struct for mapping de request params
		var RequestDTO struct {
			User struct {
				Name string
			}
		}

		//02 - make parser of params to struct
		parser.Parser(c.Request, &RequestDTO)

		//03 - Return the struct for view the result
		c.JSON(200, RequestDTO)

	})
	r.Run(":3000") // listen and serve on port:3000
}
