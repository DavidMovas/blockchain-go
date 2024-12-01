proto-gen:
	protoc -I ./node/rpc --go_out=./node/rpc --go-grpc_out=./node/rpc ./node/rpc/account.proto ./node/rpc/tx.proto ./node/rpc/block.proto

$HOST=127.0.0.1
$PORT=2233

run:
	go run . chain