// install protobuf https://github.com/protocolbuffers/protobuf
// go get  google.golang.org/grpc/cmd/protoc-gen-go-grpc
// go get google.golang.org/protobuf/cmd/protoc-gen-go

// create stubs
// protoc -I /home/lzuccarelli/Projects/protobuf --go_out=/home/lzuccarelli/Projects/protobuf /home/lzuccarelli/Projects/protobuf/internal/proto-files/service/schema-service.proto
// protoc -I /home/lzuccarelli/Projects/protobuf --go_out=/home/lzuccarelli/Projects/protobuf --go-grpc_out=/home/lzuccarelli/Projects/protobuf /home/lzuccarelli/Projects/protobuf/internal/proto-files/service/schema-service.proto

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/impl"
	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/service"
	"github.com/luigizuccarelli/simple-gRPC/pkg/validator"
	"github.com/microlib/simple"
	gRPC "google.golang.org/grpc"
)

var (
  logger *simple.Logger
)

func main() {

	if os.Getenv("LOG_LEVEL") == "" {
		logger = &simple.Logger{Level: "info"}
	} else {
		logger = &simple.Logger{Level: os.Getenv("LOG_LEVEL")}
	}
	err := validator.ValidateEnvars(logger)
	if err != nil {
		os.Exit(-1)
	}
	port, _ := strconv.Atoi(os.Getenv("GRPCSERVER_PORT"))
	l := getNetListener(port)
	gRPCServer := gRPC.NewServer()
	dataschemaServiceImpl := impl.NewDataSchemaServiceGrpcImpl()
	service.RegisterDataSchemaServiceServer(gRPCServer, dataschemaServiceImpl)
	logger.Info(fmt.Sprintf("GRPC Server starting on port %d",port))
	if err := gRPCServer.Serve(l); err != nil {
		logger.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(-1)
	}
	
}

func getNetListener(port int) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen: %v", err))
		os.Exit(-1)
	}
	return lis
}
