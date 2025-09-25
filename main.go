package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := []byte("test")
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
}
