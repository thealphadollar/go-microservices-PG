package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println(user)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("email.service"),
	)

	service.Init()

	micro.RegisterSubscriber(topic, service.Server(), new(Subscriber))

	if err := service.Run(); err != nil {
		log.Fatalf("failed to run email service, %v", err)
	}
}

func sendEmail(user *pb.User) error {
	log.Printf("sent the email to %s", user.Email)
	return nil
}
