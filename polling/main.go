package main

import (
	"log"
	"net"

	cf "poll-service/config"
	"poll-service/storage"

	pb "poll-service/genprotos"
	service "poll-service/service"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()
	em := cf.NewErrorManager()
	db, err := storage.NewPostgresStorage(config)
	em.CheckErr(err)
	defer db.PgClient.Close()

	listener, err := net.Listen("tcp", config.POLL_SERVICE_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPollServiceServer(s, service.NewPollService(db))
	pb.RegisterQuestionServiceServer(s, service.NewQuestionService(db))
	pb.RegisterUserServiceServer(s, service.NewUserService(db))
	pb.RegisterResultServiceServer(s, service.NewResultService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
