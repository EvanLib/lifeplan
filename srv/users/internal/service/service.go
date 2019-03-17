package service

import (
	"context"
	"log"

	"errors"

	"github.com/evanlib/lifeplan/srv/users/internal/database"
	users "github.com/evanlib/lifeplan/srv/users/proto"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

const (
	CollectionUsers   = "users"
	userErrorNotFound = "user with specified data not found"
)

type UsersService struct {
	database *database.Source
	auth     *AuthKeys
}

func NewUsersService(db *database.Source) *UsersService {
	auth := initAuth()
	return &UsersService{
		database: db,
		auth:     auth,
	}
}

func (us *UsersService) Create(ctx context.Context, req *users.User, res *users.UserResponse) error {

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)

	user := &users.User{
		Id:       bson.NewObjectId().Hex(),
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err = us.database.Collection(CollectionUsers).Insert(user)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (us *UsersService) Get(ctx context.Context, req *users.User, res *users.UserResponse) error {
	var user *users.User

	query := bson.M{"_id": req.Id}
	err := us.database.Collection(CollectionUsers).Find(query).One(&user)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

func (us *UsersService) GetAll(ctx context.Context, req *users.Request, res *users.UserResponse) error {
	var userlist []*users.User

	err := us.database.Collection(CollectionUsers).Find(nil).All(&userlist)
	if err != nil {
		return err
	}

	res.Users = userlist
	return nil
}

func (us *UsersService) Auth(ctx context.Context, req *users.User, resp *users.Token) error {
	var user *users.User
	query := bson.M{"email": req.Email}
	err := us.database.Collection(CollectionUsers).Find(query).One(&user)
	if err != nil {
		if err != mgo.ErrNotFound {
			log.Printf("Query to find user by id failed.", []interface{}{"err", err.Error(), "query", query})
		}

		return errors.New(userErrorNotFound)
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	// create token
	token, err := us.auth.Encode(user)
	if err != nil {
		return err
	}

	resp.Token = token
	return nil
}

func (us *UsersService) ValidateToken(ctx context.Context, req *users.Token, resp *users.Token) error {

	tokenauth, err := us.auth.Decode(req.Token)
	if err != nil {
		return err
	}

	if tokenauth.UserID == "" {
		return errors.New("invalid user")
	}
	resp.Valid = true
	return err
}
