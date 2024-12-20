package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secretKey := []byte("Sgb9F0a0sS1VpDtmwd9tpR8HS8YYReEHTzLdtG6JDEk")

	claims := jwt.MapClaims{
		"user_id":  1,
		"username": "admin",
		"roles":    []string{"ADMIN"},
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nGenerated Token:\n%s\n", signedToken)
}
