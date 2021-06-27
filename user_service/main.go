package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
)

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(255) not null unique,
		password varchar(255) not null,
		company varchar(125),
		primary key (id)
	);
`

func main() {
	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	db.MustExec(schema)

	repo, err := NewPostgresRepositoy(db)
	if err != nil {
		log.Panic(err)
	}
	tokenService := &TokenService{repo}
	service := micro.NewService(
		micro.Name("user.service"),
		micro.Version("latest"),
	)
	service.Init()

	pubsub := micro.NewPublisher("user.created", service.Client())

	if err := pb.RegisterUserServiceHandler(service.Server(), &handler{repo, tokenService, pubsub}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
