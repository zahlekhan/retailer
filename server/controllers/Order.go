package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zahlekhan/retailer/server/models"
	"net/http"
)

type CreateOrderRequest struct {
	CustomerID uint `json:"customer_id"`
	Products   []struct {
		ID       uint `json:"id"`
		Quantity uint `json:"quantity"`
	} `json:"products"`
}

func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	var err error

	_ = c.BindJSON(&req)
	cid := req.CustomerID
	var ProductIDs []uint
	var Quantities []uint
	for _, product := range req.Products {
		ProductIDs = append(ProductIDs, product.ID)
		Quantities = append(Quantities, product.Quantity)
	}
	fmt.Println("Customer ID", cid)
	o := models.Order{CustomerID: cid}
	err = models.BatchCreateOrderBookByOrderID(&o, ProductIDs, Quantities)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, o)
	}
}

//GetOrderByID ... Get the order by id
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var o models.Order
	err := models.FindOrderByID(&o, id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, o)
	}
}
