package main

import (
	"context"
	"fmt"
	//	"fmt"
	"log"
	"net/http"

	pb "inventory_SKU/grpc"

	// "inventory_SKU/models"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var(
	client pb.InventoryServiceClient
//	server         *gin.Engine
)

func main() {
	router := gin.Default()
	conn, err := grpc.Dial("localhost:5003", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client = pb.NewInventoryServiceClient(conn)


	router.POST("/api/inventory/create", CreateInventory)
	router.POST("/api/inventory/createmore", CreateMoreInventory)
  	router.POST("/api/inventory/update", UpdateInventory)

	router.Run(":3000")
}


func CreateInventory(ctx *gin.Context){
	var request pb.InventorySKU

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := client.CreateInventory(context.Background(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"value": response})
}

func CreateMoreInventory(ctx *gin.Context){
	var request []pb.InventorySKU
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(len(request))
	var responseValues []string
    for _, requests := range request {
        response, err := client.CreateInventory(context.Background(), &requests)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        responseValues = append(responseValues, response.Response)
    }

    ctx.JSON(http.StatusOK, gin.H{"values": responseValues})

}

func UpdateInventory(ctx *gin.Context){
	var request pb.UpdatedInventory
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := client.UpdateInventory(context.Background(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"value": response})
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