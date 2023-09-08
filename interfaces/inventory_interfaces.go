package interfaces

import (
	"inventory_SKU/models"

	//"go.mongodb.org/mongo-driver/mongo"
)

type IInventory interface {
	CreateInventory(user *models.Inventory_SKU) (string, error)
	//CreateMoreInventory(inventories []*models.Inventory_SKU)(string,error)
	UpdateInventory(pass * models.UpdatedInventory) (error)
	IsInStock(sku string) (bool, error)
}
