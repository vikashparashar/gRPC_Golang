# gRPC Go

## Notes


### `Windows - 

shell
protoc -Igreet/proto --go_opt=module=github.com/vikashparashar/gRPC_Golang --go_out=. --go-grpc_opt=module=github.com/vikashparashar/gRPC_Golang --go-grpc_out=. greet/proto/*.proto

protoc -Icalculator/proto --go_opt=module=github.com/vikashparashar/gRPC_Golang --go_out=. --go-grpc_opt=module=github.com/vikashparashar/gRPC_Golang --go-grpc_out=. calculator/proto/*.proto


go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
