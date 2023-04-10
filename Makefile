.PHONY: protos

protos:
	 protoc -I protos/ protos/currency.proto --go_out=. --go-grpc_out=. protos/currency.proto