package grpc_client

import (
	"log"
	"sync"

	pb "inventory_SKU/grpc"

	"google.golang.org/grpc"
)

var once sync.Once

type GrpcClient pb.InventoryServiceClient

var (
	instance GrpcClient
)

func GetGrpcClientInstance() (GrpcClient, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	once.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost:5003", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		instance = pb.NewInventoryServiceClient(conn)
	})

	return instance, conn
}
