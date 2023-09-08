package controllers

import (
	//"inventory_SKU/interfaces"
	"context"
	"fmt"
	pro "inventory_SKU/grpc"
	"inventory_SKU/interfaces"
	"inventory_SKU/models"
	//	"net/http"
	//"strings"
	// "github.com/gin-gonic/gin"
)

type RPCServer struct {
	pro.UnimplementedInventoryServiceServer
}

var (
	InventoryService interfaces.IInventory
	// UpdateService    interfaces.IUpdateInventory
	// MoreInventoryService interfaces.IMoreinventory
)

func (s *RPCServer) CreateInventory(ctx context.Context, req *pro.InventorySKU) (*pro.InventoryResponse, error) {
	dbCustomer := &models.Inventory_SKU{

		Sku: req.Sku,
		Price: models.Price{
			Base:     req.Price.Base,
			Currency: req.Price.Currency,
			Discount: req.Price.Discount,
		},
		Quantity: req.Quantity,
		Options: models.Options{
			Size: models.Size{
				H: req.Options.Size.H,
				L: req.Options.Size.L,
				W: req.Options.Size.W,
			},
			Features: req.Options.Features,
			Colors:   req.Options.Colors,
			Ruling:   req.Options.Ruling,
			Image:    req.Options.Image,
		},
	}

	_, err := InventoryService.CreateInventory(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.InventoryResponse{
			Response: "success",
		}

		return responseCustomer, nil
	}

}

func (s *RPCServer) UpdateInventory(ctx context.Context, req *pro.UpdatedInventory) (*pro.InventoryResponse, error) {

		pass := models.UpdatedInventory{
			Sku:      req.Sku,
			Quantity: req.Quantity,
		}

	err := InventoryService.UpdateInventory(&pass)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	responseCustomer := &pro.InventoryResponse{
		Response: "success",
	}
	return responseCustomer, nil
}



