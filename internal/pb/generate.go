package pb

//go:generate protoc --go_out=. --proto_path=proto common.proto
//go:generate protoc --go_out=. --proto_path=proto errors.proto

////go:generate protoc --go_out=. --proto_path=proto rpc.proto
//go:generate protoc --go_out=plugins=grpc:. --proto_path=proto rpc.proto
