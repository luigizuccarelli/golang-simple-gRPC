package impl

import (
	"context"
	"log"
	"strconv"

	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/domain"
	"github.com/luigizuccarelli/simple-gRPC/internal/gRPC/service"
)

// DataSchemaServiceGrpcImpl is a implementation of DataSchema Grpc Service.
type DataSchemaServiceGrpcImpl struct {
}

//NewDataSchemaServiceGrpcImpl returns the pointer to the implementation.
func NewDataSchemaServiceGrpcImpl() *DataSchemaServiceGrpcImpl {
	return &DataSchemaServiceGrpcImpl{}
}

//Add function implementation of gRPC Service.
func (serviceImpl *DataSchemaServiceGrpcImpl) Get(ctx context.Context, in *domain.DataSchema) (*service.GetDataSchemaResponse, error) {
	log.Println("Received request getting dataschema with id " + strconv.FormatInt(in.Id, 10))

	return &service.GetDataSchemaResponse{
		Dataschema: in,
		Error:      nil,
	}, nil
}
