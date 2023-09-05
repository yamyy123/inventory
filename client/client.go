package main

import (
	"context"
	"fmt"
	"log"

	pb "inventory_SKU/grpc"
	"inventory_SKU/models"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	/*This line establishes a connection to a gRPC server running on "localhost" at port 5001. It uses the grpc.Dial function to create a gRPC client connection.
	grpc.WithInsecure() is used here, which means that the connection is not using secure TLS/SSL. In production, it's advisable to use a secure connection.*/

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		/*This block of code checks if there was an error when attempting to establish the gRPC connection. If there was an error, it logs a fatal error message and terminates the program using log.Fatalf*/
	}

	defer conn.Close()
	/*This ensures that the connection is properly closed when the program finishes its execution.*/

	client := pb.NewInventoryServiceClient(conn)

	/* This line creates a gRPC client object named client. It uses the pb.NewCustomerServiceClient constructor to create a client that can call methods defined in the "module/netxd_customer" gRPC service.*/

	response, err := client.CreateInventory(context.Background(), &pb.InventorySKU{
		
			Sku: "YourSKUValue",
			Price: &pb.Price{
				Base:     10.0,         // Set your desired values
				Currency: "USD",        // Set your desired values
				Discount: 5.0,          // Set your desired values
			},
			Quantity: 100.0,           // Set your desired value
			Options: &pb.Options{
				Size: &pb.Size{
					H: 10.0,            // Set your desired values
					L: 20.0,            // Set your desired values
					W: 5.0,             // Set your desired values
				},
				Features: []string{"Feature1", "Feature2"}, // Set your desired values
				Colors:   []string{"Red", "Blue"},          // Set your desired values
				Ruling:   "Ruled",                          // Set your desired value
				Image:    "example.jpg",                    // Set your desired value
			},
		})
	/*This line calls the CreateCustomer method on the gRPC client client. It passes a context (context.Background()) and a pb.CustomerRequest object with some customer-related data.
	It stores the response in the response variable and any error in the err variable. */

	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Sku)
}
