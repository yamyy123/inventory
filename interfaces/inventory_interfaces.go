package interfaces

import (
	"inventory_SKU/models"

	//"go.mongodb.org/mongo-driver/mongo"
)

type IInventory interface {
	CreateInventory(user *models.Inventory_SKU) (string, error)
	//CreateMoreInventory(inventories []*models.Inventory_SKU)(string,error)
	UpdatedInventory(pass * models.UpdatedInventory) (error)
	IsInStock(sku string) (bool, error)
}
