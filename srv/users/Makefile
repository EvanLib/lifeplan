
protoc:
	for f in proto/*.proto; do \
		echo compiled: $$f; \
	done
	protoc \
	--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
	--micro_out=. --go_out=plugins=grpc:. proto/users.proto \
	