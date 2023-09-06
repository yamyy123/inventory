package interfaces

import "inventory_SKU/models"

type IUpdateInventory interface {
	UpdateInventory(pass * models.UpdatedInventory) (error)
	IsInStock(sku string) (bool, error)
}
