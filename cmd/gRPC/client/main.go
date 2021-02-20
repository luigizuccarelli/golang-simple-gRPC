package main

import (
	"context"
	"fmt"
	"os"

	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/domain"
	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/service"

	gRPC "google.golang.org/grpc"
)

func main() {
	serverAddress := os.Args[1] + ":7000"

	conn, e := gRPC.Dial(serverAddress, gRPC.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewDataSchemaServiceClient(conn)

	for i := range [10]int{} {
		dataschemaModel := domain.DataSchema{
			Id:      int64(i + 1),
			Name:    string("Grpc-Demo"),
			Status:  string("OK"),
			Payload: string("{\"message\":\"dude this sh*t is working\"}"),
		}

		if responseMessage, e := client.Get(context.Background(), &dataschemaModel); e != nil {
			panic(fmt.Sprintf("Was not able to get Record %v", e))
		} else {
			fmt.Println("DataSchema reedge..")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
