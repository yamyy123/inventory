package services

import (
	"context"
	"errors"
	"fmt"
	"inventory_SKU/interfaces"
	"inventory_SKU/models"

	//"inventory_SKU/interfaces"
	//"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryService struct {
	client              *mongo.Client
	InventoryCollection *mongo.Collection
	ctx                 context.Context
}

type UpdateInventoryService struct {
	collection *mongo.Collection
}


func NewUpdatedInventoryServiceInit(collection *mongo.Collection) interfaces.IUpdateInventory {
	return &UpdateInventoryService{
		collection: collection,
	}

}

func NewInventoryServiceInit(client *mongo.Client, collection *mongo.Collection, ctx context.Context) interfaces.IInventory {
	return &InventoryService{client, collection, ctx}

}
func (b *InventoryService) CreateInventory(user *models.Inventory_SKU) (string, error) {

	result, err := b.InventoryCollection.InsertOne(b.ctx, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return "success", nil

}

// IsInStock checks if an item with the given SKU is in stock.
func (s *UpdateInventoryService) IsInStock(sku string) (bool, error) {
	filter := bson.M{"sku": sku}
	var item models.Inventory_SKU

	err := s.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Item not found, consider it out of stock
		}
		return false, err
	}

	return item.Quantity > 0, nil // In stock if quantity is greater than zero
}

// UpdatedInventory updates the inventory for an item.
func (s *UpdateInventoryService) UpdateInventory(pass *models.UpdatedInventory) error {
	filter := bson.M{"sku": pass.Sku}
	fmt.Println(filter)
	var item models.Inventory_SKU

	err :=s.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("item with SKU %s not found", pass.Sku)
		}
		return err
	}

	inStock, err := s.IsInStock(pass.Sku)
	if err != nil {
		return err
	}

	if !inStock {
		return fmt.Errorf("item with SKU %s is not in stock", pass.Sku)
	}

	if item.Quantity < pass.Quantity {
		return fmt.Errorf("not enough quantity in stock, only %v available", item.Quantity)
	}

	update := bson.M{"$inc": bson.M{"quantity": -pass.Quantity}}

	_, err = s.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
