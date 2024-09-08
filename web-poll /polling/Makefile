run:
	clear; go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ protos/*.proto
	
proto-gen0:
	protoc --go_out=./ --go-grpc_out=./ --experimental_allow_proto3_optional protos/*.proto

migrate_up:
	migrate -path migrations -database postgres://postgres:root@localhost:5432/polling_db?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:root@localhost:5432/polling_db?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:root@localhost:5432/polling_db?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table


swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1

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
