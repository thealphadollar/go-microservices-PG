package main

import (
	"context"
	"errors"

	pb "github.com/thealphadollar/go-microservices-PG/user_service/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type handler struct {
	repo         Repository
	tokenService authable
}

func (h *handler) Create(ctx context.Context, user *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	if err := h.repo.Create(ctx, MarshalUser(user)); err != nil {
		return err
	}
	user.Password = ""
	res.User = user
	res.Errors = nil

	return nil
}

func (h *handler) Get(ctx context.Context, user *pb.User, res *pb.Response) error {
	result, err := h.repo.Get(ctx, user.Id)
	if err != nil {
		return err
	}
	res.User = UnmarshalUser(result)
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	result, err := h.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Users = UnmarshalUserCollection(result)
	return nil
}

func (h *handler) Auth(ctx context.Context, user *pb.User, res *pb.Token) error {
	result, err := h.repo.GetByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return err
	}
	token, err := h.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (h *handler) ValidateToken(ctx context.Context, token *pb.Token, res *pb.Token) error {
	claims, err := h.tokenService.Decode(token.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user token")
	}
	res.Token = token.Token
	res.Valid = true
	return nil
}
