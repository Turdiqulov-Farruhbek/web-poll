run:
	clear; go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ protos/*.proto
	
migrate_up:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/polling_db?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/polling_db?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://mrbek:QodirovCoder@localhost:5432/polling_db?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1