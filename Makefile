.PHONY: proto;

protos:
	protoc -I order/ proto/order.proto --go_out=./order/proto --go-grpc_out=./order/proto;
	protoc -I user/ proto/user.proto --go_out=./user/proto --go-grpc_out=./user/proto;