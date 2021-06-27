package main

import (
	"context"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
)

const (
	name     = "Shivam"
	email    = "s-kumar3@mercari.com"
	password = "somepass"
	company  = "Mercari"
)

func main() {
	service := micro.NewService(
		micro.Name("user.cli"),
		micro.Version("latest"),
	)

	service.Init()

	client := pb.NewUserService("user.service", service.Client())
	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not get all users: %v", err)
	}

	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResp, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("failed to login: %v", err)
	}

	log.Printf("Your access token: %s", authResp.Token)

	_, err2 := client.ValidateToken(context.Background(), &pb.Token{
		Token: authResp.Token,
	})

	if err2 != nil {
		log.Fatalf("failed to validate, %v", err2)
	}

	os.Exit(0)
}
