# Simple protobuf golang services (client and server)

A simple golang protobuf example.

## Install protobuf
- navigate to https://github.com/protocolbuffers/protobuf - releases (download and install for your os)
- ```go get  google.golang.org/grpc/cmd/protoc-gen-go-grpc```
- ```go get google.golang.org/protobuf/cmd/protoc-gen-go```

- create project directory structure


```
projects/protobuf (top level)
  - cmd
    - server
      -- main.go
    - client
      -- main.go
  - internal
    - gRPC
      - domain
        -- (stub schema.pg.go will be created here)
      - impl
        -- schema-service.go
      - service
        -- (stub schema-service.pg.go will be created here)
  - proto-files
    - domain
      -- schema.proto
    - service
      -- schema-service.proto
  - pkg
      
```

- define your schema in the domain/schema.proto file (example)


```
syntax = "proto3";

package domain;

option go_package = "internal/gRPC/domain";

message DataSchema {
  int64 id = 1;
  string name = 2;
  string status = 3;
  string payload = 4;
}

```

- define the service data in the file service/schema-service.proto (example)

```

syntax = "proto3";

package service;

option go_package = "internal/gRPC/service";

import "internal/proto-files/domain/schema.proto";

service DataSchemaService {
	rpc get (domain.DataSchema) returns (GetDataSchemaResponse);
}
 
message GetDataSchemaResponse {
	domain.DataSchema dataschema = 1;
	Error error = 2;
}

message Error {
	string code = 1;
	string message = 2;
}
```

- now generate the stubs

```
// create stubs
$ protoc -I /home/lzuccarelli/Projects/protobuf --go_out=/home/lzuccarelli/Projects/protobuf /home/lzuccarelli/Projects/protobuf/internal/proto-files/service/schema-service.proto

$ protoc -I /home/lzuccarelli/Projects/protobuf --go_out=/home/lzuccarelli/Projects/protobuf --go-grpc_out=/home/lzuccarelli/Projects/protobuf /home/lzuccarelli/Projects/protobuf/internal/proto-files/service/schema-service.proto
```

- update client.go and server.go

- build using make

## Testing

- scp client to pi cluster (ip 192.168.1.14)
- start server
- start client on remote pi server


