package pb

//go:generate protoc --go_out=. --proto_path=proto common.proto
//go:generate protoc --go_out=. --proto_path=proto errors.proto
//go:generate protoc --go_out=plugins=grpc:. --proto_path=proto logic.proto
//go:generate protoc --go_out=plugins=grpc:. --proto_path=proto api.proto
