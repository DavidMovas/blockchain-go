proto-gen:
	protoc -I ./node/rpc --go_out=./node/rpc --go-grpc_out=./node/rpc ./node/rpc/account.proto