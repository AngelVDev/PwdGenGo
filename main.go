package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generatePassword(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	password := base64.URLEncoding.EncodeToString(buffer)
	return password[:length], nil
}

func main() {
	password, err := generatePassword(12)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}

	fmt.Println("Generated password:", password)
}
