package service

import (
	"context"
	"log"

	"github.com/evanlib/lifeplan/srv/users/internal/database"
	users "github.com/evanlib/lifeplan/srv/users/proto"
	"github.com/globalsign/mgo/bson"
)

const (
	CollectionUsers = "users"
)

type UsersService struct {
	database *database.Source
}

func NewUsersService(db *database.Source) *UsersService {
	return &UsersService{
		database: db,
	}
}

func (UsersService *UsersService) Create(ctx context.Context, req *users.User, res *users.UserResponse) error {

	log.Println(req)
	// marsahl user
	userMarshalled, err := bson.Marshal(req)
	if err != nil {
		return err
	}

	err = UsersService.database.Collection(CollectionUsers).Insert(userMarshalled)
	if err != nil {
		return err
	}
	return nil
}

func (UsersService *UsersService) Get(ctx context.Context, req *users.User, res *users.UserResponse) error {
	return nil
}

func (UsersService *UsersService) GetAll(ctx context.Context, req *users.Request, res *users.UserResponse) error {
	return nil
}

func (UsersService *UsersService) Auth(ctx context.Context, req *users.User, resp *users.Token) error {
	log.Println("[USERS] Log in with: ", req.Email, req.Password)
	return nil
}

func (UsersService *UsersService) ValidateToken(ctx context.Context, req *users.Token, resp *users.Token) error {
	return nil
}
