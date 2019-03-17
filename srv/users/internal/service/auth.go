package service

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	pb "github.com/evanlib/lifeplan/srv/users/proto"
)

const (
	privKeyPath = "internal/authentication/app.rsa"
	pubKeyPath  = "internal/authentication/app.rsa.pub"
)

type Claims struct {
	UserID    string
	UserName  string
	UserEmail string
	*jwt.StandardClaims
}

type AuthKeys struct {
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
}

func initAuth() *AuthKeys {

	authkeys := &AuthKeys{}

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

	authkeys.verifyKey = verifyKey
	authkeys.signKey = signKey

	return authkeys
}

func (ak *AuthKeys) Encode(user *pb.User) (string, error) {
	// create signer for rsa256
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// create claims
	t.Claims = &Claims{
		user.Id,
		user.Username,
		user.Email,
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create Token
	return t.SignedString(ak.signKey)
}

func (ak *AuthKeys) Decode(decodetoken string) (*Claims, error) {
	//Parse Token
	token, err := jwt.ParseWithClaims(decodetoken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return ak.verifyKey, nil
	})

	// validate the token and return claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
