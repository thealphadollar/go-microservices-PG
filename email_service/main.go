package main

import (
	"encoding/json"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	_ "github.com/micro/go-plugins/broker/nats/v2"
	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
)

const topic = "user.created"

func main() {
	service := micro.NewService(
		micro.Name("email.service"),
	)

	service.Init()

	pubsub := service.Server().Options().Broker

	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	_, err := pubsub.Subscribe(topic, func(p broker.Event) error {
		var user *pb.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil {
			log.Fatal(err)
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Fatalf("failed to subscribe to pubsub: %v", err)
	}

	if err := service.Run(); err != nil {
		log.Fatalf("failed to run email service, %v", err)
	}
}

func sendEmail(user *pb.User) error {
	log.Printf("sent the email to %s", user.Email)
	return nil
}
