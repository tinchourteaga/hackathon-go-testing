package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga-ml/desafio-go-testing-martin-urteaga/cmd/router"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")

}
