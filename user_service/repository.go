package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
	_ "github.com/lib/pq"
)

type User struct {
	Id string `sql:"id"`
	Name string `sql:"name"`
	Company string `sql:"company"`
	Password string `sql:"password"`
	Email string `sql:"email"`
}

type Repository interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type PostgresqlRepo struct {
	db *sqlx.DB
}

func NewPostgresRepositoy (db *sqlx.DB) (*PostgresqlRepo, error) {
	return &PostgresqlRepo{db}, nil
}

func MarshalUser(user *pb.User) *User {
	return &User{
		Id: user.Id,
		Company: user.Company,
		Name: user.Name,
		Password: user.Password,
		Email: user.Email,
	}
}

func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id: user.Id,
		Company: user.Company,
		Name: user.Name,
		Password: user.Password,
		Email: user.Email,
	}
}

func UnmarshalUserCollection(users []*User) []*pb.User {
	u := make([]*pb.User, len(users))
	for _, v := range users {
		u = append(u, UnmarshalUser(v))
	}
	return u
}

func (r *PostgresqlRepo) Create (ctx context.Context, user *User) error {
	user.Id = uuid.NewV4().String()
	query := "insert into users (id, name, email, company, password) values ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, user.Id, user.Name, user.Email, user.Company, user.Password)
	return err
}

func (r *PostgresqlRepo) Get (ctx context.Context, id string) (*User, error) {
	var user *User
	query := "select * from users where id = $1"
	if err := r.db.GetContext(ctx, &user, query, id); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresqlRepo) GetAll (ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err:= r.db.GetContext(ctx, users, "select * from users"); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresqlRepo) GetByEmail (ctx context.Context, email string) (*User, error) {
	var user *User
	query := "select * from users where email = $1"
	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}
	return user, nil
}