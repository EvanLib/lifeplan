package main

import (
	"log"
	"os"

	pb "github.com/evanlib/lifeplan/srv/users/proto"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

func main() {

	cmd.Init()

	// create new user client
	client := pb.NewUsersService("go.micro.src.users", microclient.DefaultClient)

	username := "Evan"
	email := "evangrayson95@gmail.com"
	password := "lol626465"

	log.Println(username, email, password)

	r, err := client.Create(context.TODO(), &pb.User{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s \n", r.User.Id)

	u, err := client.Get(context.TODO(), &pb.User{
		Id: r.User.Id,
	})
	if err != nil {
		log.Fatalf("Could not get: %v", err)
	}
	log.Printf("Get: %s \n", u.User.Username)

	user, err := client.GetAll(context.TODO(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not getall: %v", err)
	}
	for _, usern := range user.Users {
		log.Printf("Username: %s \n", usern.Username)
	}

	// check auth
	userauth := &pb.User{
		Email:    email,
		Password: password,
	}
	token, err := client.Auth(context.TODO(), userauth)
	if err != nil {
		log.Printf("Problem with auth login: %s", err)
	}
	log.Printf("User auth returned token: %s", token.Token)

	authTry, err := client.ValidateToken(context.TODO(), &pb.Token{Token: token.Token})
	if err != nil {
		log.Printf("Problem on validating token: %s", err)
	}
	log.Printf("Token Validated returned token: %s", authTry)
	os.Exit(0)
}
