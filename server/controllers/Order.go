package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zahlekhan/retailer/server/models"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var o models.Order
	var err error
	_ = c.BindJSON(&o)
	err = models.CreateOrder(&o)
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
