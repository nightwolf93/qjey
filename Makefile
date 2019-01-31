run:
	go build -o ./bin/qjey
	cp .env ./bin/.env
	cd ./bin && ./qjey
