run:
	clear; go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ protos/*.proto
	
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



# doc-continer-all-del:
# 	docker rm $(docker ps -aq)
# 	docker rmi $(docker ps images)

# doc-network-all:
# 	docker network ls


# doc-query:
#   docker rm ...
#   docker rmi -f ...
#   docker container ps
# 	docker ps
# 	docker image rmi auth_service
# 	docker container ls
# 	docker container prune
# 	docker container rm libCont
# 	docker build -t lib_service .
# 	docker network ls
# 	docker images
# 	docker exec -it post-cont psql -U postgres -d library_db
# 	docker system prune -af
#  	docker volume prune -af