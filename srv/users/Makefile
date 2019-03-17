
protoc:
	protoc \
	--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
	--micro_out=. --go_out=plugins=micro:. proto/users.proto \
	
	protoc-go-inject-tag -XXX_skip=bson -input=./proto/users.pb.go