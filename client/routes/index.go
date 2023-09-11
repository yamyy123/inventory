package routes

import (
	"inventory_SKU/client/controller"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) {
	router.POST("/api/inventory/create", controller.CreateInventory)
	router.POST("/api/inventory/createmore", controller.CreateMoreInventory)
	router.POST("/api/inventory/update", controller.UpdatedInventory)

}
