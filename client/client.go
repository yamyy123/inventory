package main

import (

	client "inventory_SKU/client/grpc_client"

	"inventory_SKU/client/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, conn := client.GetGrpcClientInstance()
	defer conn.Close()

	r := gin.Default()
	routes.AppRoutes(r)
	r.Run(":3000")
}

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
//-----------------------------------------------------------------------------------------
// response, err := client.UpdateInventory(context.Background(), &pb.UpdatedInventory{
// 	Sku:      "SKU002",
// 	Quantity: 4.0,
// })
// if err != nil {
// 	log.Fatalf("Failed to call UpdatedInventory: %v", err)
// }

// fmt.Printf("Response: %s\n", response.Response)
// ---------------------------------------------------------------------------------------------

// response, err := client.CreateMoreInventory(context.Background(), &pb.MoreInventories{
// 	InventorySKU: []*pb.InventorySKU{
// 		{
// 			Sku: "SKU005",
// 			Price: &pb.Price{
// 				Base:     1.0,
// 				Currency: "inr",
// 				Discount: 3.0,
// 			},
// 			Quantity: 100.0,
// 			Options: &pb.Options{
// 				Size: &pb.Size{
// 					H: 10.0,
// 					L: 5.0,
// 					W: 2.0,
// 				},
// 				Features: []string{"Feature 1", "Feature 2"},
// 				Colors:   []string{"Red", "Blue"},
// 				Ruling:   "Ruled",
// 				Image:    "image.jpg",
// 			},
// 		},
// 		{
//             Sku: "SKU006",
// 			Price: &pb.Price{
// 				Base:     1.0,
// 				Currency: "inr",
// 				Discount: 3.0,
// 			},
// 			Quantity: 100.0,
// 			Options: &pb.Options{
// 				Size: &pb.Size{
// 					H: 10.0,
// 					L: 5.0,
// 					W: 2.0,
// 				},
// 				Features: []string{"Feature 1", "Feature 2"},
// 				Colors:   []string{"Red", "Blue"},
// 				Ruling:   "Ruled",
// 				Image:    "image.jpg",
// 			},
// 		},
// 	},
// })

// if err != nil {
// 	log.Fatalf("Failed to call SayHello: %v", err)
// }

// fmt.Printf("Response: %s\n", response)
