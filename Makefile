gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=./pb proto/*.proto --go_out=pb

clean:
	rm pb/*.go
