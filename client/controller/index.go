package controller

import (
	"context"
	"fmt"
	grpcclient "inventory_SKU/client/grpc_client"
	pb "inventory_SKU/grpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInventory(ctx *gin.Context) {

	client, _ := grpcclient.GetGrpcClientInstance()

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

func CreateMoreInventory(ctx *gin.Context) {

	client, _ := grpcclient.GetGrpcClientInstance()
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

func UpdatedInventory(ctx *gin.Context) {
	client, _ := grpcclient.GetGrpcClientInstance()
	var request pb.UpdateInventory
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := client.UpdatedInventory(context.Background(), &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"value": response})

}
