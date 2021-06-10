package main

import (
	"context"
	"fmt"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
)

func CreateUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("user.service", service.Client())
	res, err := client.Create(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println("Response: ", res.User.Id, res.User.Company)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "your name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "your email",
			},
			&cli.StringFlag{
				Name:  "company",
				Usage: "your company",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "your password",
			},
		),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			ctx := context.Background()
			user := &pb.User{
				Name:     name,
				Email:    email,
				Company:  company,
				Password: password,
			}

			if err := CreateUser(ctx, service, user); err != nil {
				fmt.Printf("failed to create user: %v", err)
				return err
			}
			return nil
		}),
	)
}
