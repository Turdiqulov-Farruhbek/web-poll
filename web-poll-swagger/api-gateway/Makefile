run:
	clear;
	go mod tidy;
	go run main.go;
migrate_up:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/delivery_auth -verbose up

migrate_down:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/delivery_auth -verbose down

migrate_force:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/delivery_auth -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ protos/*.proto

proto-gen0:
	protoc --go_out=./ --go-grpc_out=./ --experimental_allow_proto3_optional protos/*.proto


ps:
	docker ps -a

pg:
	docker exec -it $(c) psql -U postgres -d polling_db

down:
	docker compose down
	
up:
	docker compose up

build:
	docker compose up --build