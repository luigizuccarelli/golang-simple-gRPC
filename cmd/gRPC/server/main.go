// install protobuf https://github.com/protocolbuffers/protobuf
// go get  google.golang.org/grpc/cmd/protoc-gen-go-grpc
// go get google.golang.org/protobuf/cmd/protoc-gen-go

// create stubs
// protoc -I /home/lzuccarelli/Projects/protobuf --go_out=/home/lzuccarelli/Projects/protobuf /home/lzuccarelli/Projects/protobuf/internal/proto-files/service/schema-service.proto
// protoc -I /home/lzuccarelli/Projects/protobuf --go_out=/home/lzuccarelli/Projects/protobuf --go-grpc_out=/home/lzuccarelli/Projects/protobuf /home/lzuccarelli/Projects/protobuf/internal/proto-files/service/schema-service.proto

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/impl"
	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/service"
	gRPC "google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(7000)
	gRPCServer := gRPC.NewServer()

	dataschemaServiceImpl := impl.NewDataSchemaServiceGrpcImpl()
	service.RegisterDataSchemaServiceServer(gRPCServer, dataschemaServiceImpl)

	// start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
