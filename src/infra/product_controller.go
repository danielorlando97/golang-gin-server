package infra

import (
	"mongo-store/src/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouteProduct(r *gin.Engine) {
	r.GET("/product", getProduct)
	r.GET("/product/:id", getProductByID)
	r.POST("/product", postProduct)
}

var products = []domain.Product{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getProduct(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func postProduct(c *gin.Context) {
	var newAlbum domain.Product

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	products = append(products, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
