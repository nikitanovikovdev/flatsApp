package middlewear

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	auth "github.com/nikitanovikovdev/flatsApp-users/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func IsAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			fmt.Fprintln(w, "token is not expected")
		}
		tokenString := c.Value

		_, err = ParseToken(tokenString)
		if err != nil {
			fmt.Fprintln(w, "not valid token")
			return
		}
		endpoint.ServeHTTP(w, r)
		if err != nil {
			fmt.Fprintln(w, "failed to authorized")
			return
		}
	})
}

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username" bson:"username"`
}

func ParseToken(tkn string) (string, error) {
	if err := initConfig(); err != nil {
		return "", fmt.Errorf("error connection to config: %v", err)
	}

	token, err := jwt.ParseWithClaims(tkn, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(GetKey()), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", err
}

func GetKey() string {
	conn, err := grpc.Dial(":8040", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatal(err)
	}

	c := auth.NewAuthClient(conn)

	key, err := c.ReturnSignKey(context.Background(), new(auth.Empty))
	if err != nil {
		log.Fatal(err)
	}
	return key.GetSigningKey()
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}
