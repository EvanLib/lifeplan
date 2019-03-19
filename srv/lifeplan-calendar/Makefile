protoc:
	
	for f in proto/*.proto; do \
		protoc \
		--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
		--micro_out=. --go_out=plugins=grpc:. $$f; \
		echo compiled: $$f; \
	done

	protoc-go-inject-tag -XXX_skip=bson -input=./proto/calendar.pb.go
	protoc-go-inject-tag -XXX_skip=bson -input=./proto/events.pb.go

protocal:
	protoc \
	--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
	--micro_out=. --go_out=plugins=micro:. proto/calendar.proto \
	
	protoc-go-inject-tag -XXX_skip=bson -input=./proto/calendar.pb.go
run:
	docker-compose build
	docker-compose up