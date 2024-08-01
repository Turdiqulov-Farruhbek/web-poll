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