package main

import (
	"context"
	"fmt"
	"time"

	pb "SDCS/json_cache"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func startClient() {
	var err error

	conn[0], err = grpc.Dial(address[2], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("fail to dial: %v", err)
	}
	fmt.Println("Start client for", address[2])

	conn[1], err = grpc.Dial(address[3], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("fail to dial: %v", err)
	}
	fmt.Println("Start client for", address[3])

	client[0] = pb.NewCacheClient(conn[0])
	client[1] = pb.NewCacheClient(conn[1])
}

// gRpc客户端Get request
func GetCache(client pb.CacheClient, req *pb.GetRequest) *pb.GetResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.Get(ctx, req)
	if err != nil {
		fmt.Println("client.Get failed.")
		return nil
	}
	return response
}

// gRpc客户端Post request
func PostCache(client pb.CacheClient, req *pb.PostRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Post(ctx, req)
	if err != nil {
		fmt.Println("client.Post failed.")
	}
}

// gRpc客户端Delete request
func DeleteCache(client pb.CacheClient, req *pb.DeleteRequest) int32 {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	response, err := client.Delete(ctx, req)
	if err != nil {
		fmt.Println("client.Delete failed.")
	}
	num := response.Num
	return num
}
