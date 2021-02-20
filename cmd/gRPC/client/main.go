package main

import (
	"context"
	"fmt"
	"os"

	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/domain"
	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/service"
	"github.com/luigizuccarelli/simple-gRPC/pkg/validator"

	"github.com/microlib/simple"
	gRPC "google.golang.org/grpc"
)

func main() {
	var logger *simple.Logger
	if os.Getenv("LOG_LEVEL") == "" {
		logger = &simple.Logger{Level: "info"}
	} else {
		logger = &simple.Logger{Level: os.Getenv("LOG_LEVEL")}
	}
	err := validator.ValidateEnvars(logger)
	if err != nil {
		os.Exit(-1)
	}

	s := os.Getenv("GRPCSERVER_HOST") + ":" + os.Getenv("GRPCSERVER_PORT")
	conn, e := gRPC.Dial(s, gRPC.WithInsecure())
	if e != nil {
		os.Exit(-1)
	}
	defer conn.Close()

	c := service.NewDataSchemaServiceClient(conn)
	// for simple testing this will be removed in production
	for i := range [10]int{} {
		dataschemaModel := domain.DataSchema{
			Id:      int64(i + 1),
			Name:    string("Grpc-Demo"),
			Status:  string("OK"),
			Payload: string("{\"message\":\"dude this sh*t is working\"}"),
		}
		if responseMessage, e := c.Get(context.Background(), &dataschemaModel); e != nil {
			logger.Error(fmt.Sprintf("Response from server %v", e))
			os.Exit(-1)
		} else {
			logger.Debug("DataSchema from server (GET rpc)")
			logger.Info(fmt.Sprintf("Data : %v", responseMessage))
		}
	}
}
