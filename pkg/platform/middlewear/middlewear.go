package middlewear

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func IsAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b := r.Header["Token"]
		if b == nil {
			fmt.Fprintln(w, "token wasn't found")
			return
		}

		_, err := ParseToken(b[0])
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
		return []byte(viper.GetString("keys.signing_key")), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		return claims.Username, nil
	}

	return "", err
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	return viper.ReadInConfig()
}