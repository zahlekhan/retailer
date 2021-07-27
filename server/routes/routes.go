package Routes

import (
	"github.com/gin-gonic/gin"
	Controllers "github.com/zahlekhan/retailer/server/controllers"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	product := r.Group("/product")
	{
		product.GET("/", Controllers.GetProducts)
		product.POST("/", Controllers.CreateProduct)
		product.GET("/:id", Controllers.GetProductByID)
		product.PUT("/:id", Controllers.UpdateProduct)
	}
	order := r.Group("/order")
	{
		order.GET("/:id", Controllers.GetOrderByID)
		order.POST("/", Controllers.CreateOrder)
	}
	return r
}
