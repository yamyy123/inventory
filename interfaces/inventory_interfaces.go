package interfaces

import (
	"inventory_SKU/models"

	//"go.mongodb.org/mongo-driver/mongo"
)

type IInventory interface {
	CreateCustomer(user *models.Inventory_SKU) (string, error)
	
}
