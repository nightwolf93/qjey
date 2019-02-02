run: build-proto
	go build -o ./bin/qjey
	cp .env ./bin/.env
	cd ./bin && ./qjey

build-proto:
	protoc qjeyserver/qjeyserverprob/qjeyserver.proto --go_out=plugins=grpc:.

test-grpc:
