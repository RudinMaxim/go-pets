package main

import (
	"github.com/gin-gonic/gin"
	"pets-server/Cat"
	"pets-server/Dog"
)

func main() {

	r := gin.Default()

	Cat.CatRouter(r)

	Dog.DogRouter(r)

	r.Run("0.0.0.0:8888")
}
