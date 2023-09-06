package main

import (
	"context"
	"fmt"
	"log"

	pb "inventory_SKU/grpc"
	// "inventory_SKU/models"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewInventoryServiceClient(conn)

	// response, err := client.CreateInventory(context.Background(), &pb.InventorySKU{

	// 		Sku:      "SKU002",
	// 		Price:     &pb.Price{
	// 			Base: 1.0 ,
	// 			Currency: "inr",
	// 			Discount: 3.0,
	// 		},
	// 		Quantity: 100.0,
	// 		Options: &pb.Options{
	// 			Size: &pb.Size{
	// 				H: 10.0,
	// 				L:  5.0,
	// 				W:  2.0,
	// 				},
	// 			Features: []string{"Feature 1", "Feature 2"},
	// 			Colors:   []string{"Red", "Blue"},
	// 			Ruling:   "Ruled",
	// 			Image:    "image.jpg",
	// 		},
	// 	})

	// if err != nil {
	// 	log.Fatalf("Failed to call SayHello: %v", err)
	// }

	// fmt.Printf("Response: %s\n", response)

	response, err := client.UpdateInventory(context.Background(), &pb.UpdatedInventory{
		Sku:      "SKU002",
		Quantity: 4.0,
	})
	if err != nil {
		log.Fatalf("Failed to call UpdatedInventory: %v", err)
	}
	
	// Assuming the response is a string
	fmt.Printf("Response: %s\n", response.Response)
}
