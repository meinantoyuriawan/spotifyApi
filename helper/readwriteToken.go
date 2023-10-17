package helper

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func WriteToken(AccToken string) {
	f, err := os.Create("token.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	storeToken(f, AccToken)

}

func ReadToken() string {
	Token, err := os.ReadFile("token.txt")

	if err != nil {
		return ""
	}

	return string(Token)
}

func WriteRefreshToken(RefreshToken string) {
	f, err := os.Create("refreshtoken.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	storeToken(f, RefreshToken)
}

func ReadRefreshToken() string {
	Token, err := os.ReadFile("refreshtoken.txt")

	if err != nil {
		return ""
	}

	return string(Token)
}

func GetClientID() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("some error occured")
	}
	return os.Getenv("CLIENT_ID")
}

func GetClientSecret() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("some error occured")
	}
	return os.Getenv("CLIENT_SECRET")
}

func storeToken(f *os.File, Token string) {

	_, err := f.WriteString(Token)

	if err != nil {
		panic(err)
	}

	f.Sync()
}
