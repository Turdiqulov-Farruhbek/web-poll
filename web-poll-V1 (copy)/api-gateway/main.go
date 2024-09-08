package main

import (
	"auth-service/api"
	"fmt"

	cf "auth-service/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()

	PollConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", config.POLL_SERVICE_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	em.CheckErr(err)
	defer PollConn.Close()

	roter := api.NewRouter(PollConn)
	fmt.Println("Server is running on port ", config.GATEWAY_SERVICE_PORT)
	if err := roter.Run(config.GATEWAY_SERVICE_PORT); err != nil {
		panic(err)
	}
}
