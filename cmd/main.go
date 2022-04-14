package main

import (
	"mongo-store/src/infra"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	infra.CreateRouteProduct(router)

	router.Run("localhost:9080")
}
