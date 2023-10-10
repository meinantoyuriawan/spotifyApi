package helper

import (
	"os"
)

func WriteToken(AccToken string) {
	f, err := os.Create("token.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.WriteString(AccToken)

	if err != nil {
		panic(err)
	}

	f.Sync()
}

func ReadToken() string {
	Token, err := os.ReadFile("token.txt")

	if err != nil {
		return ""
	}

	return string(Token)
}
