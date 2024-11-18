.PHONY: proto;

protos:
	protoc -I order/ proto/order.proto --go_out=./order/proto --go-grpc_out=./order/proto;
	protoc -I user/ proto/user.proto --go_out=./user/proto --go-grpc_out=./user/proto;
order_app:
	go run order/main.go
user_app:
	go run user/main.go
tests:
	go test -v order/server/order_test.go
	go test -v user/server/