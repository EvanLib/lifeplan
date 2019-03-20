protoc:
	protoc -I=. -I=$(GOPATH)/src/github.com/gogo/protobuf -I=$(GOPATH)/src \
	--micro_out=. --gogo_out=plugins=micro:. proto/calendar.proto
	
	# Add BSON SKIP for XXX type
	protoc-go-inject-tag -XXX_skip=bson -input=./proto/calendar.pb.go
run:
	docker-compose build
	docker-compose up