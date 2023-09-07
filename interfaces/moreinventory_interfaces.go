package interfaces

import (
	"inventory_SKU/models"
)

type IMoreinventory interface{
       CreateMoreInventory(inventories []*models.Inventory_SKU)(string,error)
}