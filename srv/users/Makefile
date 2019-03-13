
protoc:
	protoc \
	--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
	--micro_out=. --go_out=plugins=grpc:. proto/users.proto \
	
	protoc-go-inject-tag -input=./proto/users.pb.go
	