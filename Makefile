.PHONY: protos;

protos:
	protoc -I proto/ proto/order.proto --go_out=. --go-grpc_out=.;
	protoc -I proto/ proto/user.proto --go_out=. --go-grpc_out=.;