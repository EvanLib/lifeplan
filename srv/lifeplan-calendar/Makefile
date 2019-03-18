protoc:
	protoc \
	--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
	--micro_out=. --go_out=plugins=micro:. proto/calendar.proto \
	
	protoc-go-inject-tag -XXX_skip=bson -input=./proto/calendar.pb.go