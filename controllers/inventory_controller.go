package controllers

import (
	//"inventory_SKU/interfaces"
	"context"
	pro "inventory_SKU/grpc"
	"inventory_SKU/interfaces"
	"inventory_SKU/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RPCServer struct{
	pro.UnimplementedInventoryServiceServer
}

var (
	InventoryService interfaces.IInventory
)

func (s *RPCServer) CreateInventory(ctx context.Context, req *pro.InventorySKU)(*pro.InventoryResponse) {
	dbCustomer := &models.Inventory_SKU{
		Sku: req.Sku,
		Price: models.Price_type{
			Base:     req.Price.Base,         // Set your desired values
			Currency: req.Price.Currency,        // Set your desired values
			Discount: 5.0,          // Set your desired values
		},
		Quantity: 100.0,           // Set your desired value
		Options: models.Options_type{
			Size: models.Size_type{
				H: 10.0,            // Set your desired values
				L: 20.0,            // Set your desired values
				W: 5.0,             // Set your desired values
			},
			Features: []string{"Feature1", "Feature2"}, // Set your desired values
			Colors:   []string{"Red", "Blue"},          // Set your desired values
			Ruling:   "Ruled",                         // Set your desired value
			Image:    "example.jpg",                   // Set your desired value
		},
	}
	
	}



