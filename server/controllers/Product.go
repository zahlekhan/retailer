package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zahlekhan/retailer/server/models"
	"net/http"
)

//GetProducts ... Get all products
func GetProducts(c *gin.Context) {
	var p []models.Product
	err := models.FindAllProducts(&p)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

//CreateProduct ... Create product
func CreateProduct(c *gin.Context) {
	var p models.Product
	var err error
	_ = c.BindJSON(&p)
	err = models.CreateProduct(&p)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

//GetProductByID ... Get the product by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var p models.Product
	err := models.FindProductByID(&p, id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

//UpdateProduct ... Update the product information
// TODO : Return updated value using transactions
func UpdateProduct(c *gin.Context) {
	var p models.Product
	id := c.Params.ByName("id")
	price := c.Params.ByName("price")
	qty := c.Params.ByName("quantity")
	err := models.UpdateProduct(id, price, qty)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.JSON(http.StatusOK, p)
	}
}
