package helper

import (
	"bufio"
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

func getClientConfig() []string {
	filePath := "config.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		return nil
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func GetClientID() string {
	clientConfig := getClientConfig()
	return clientConfig[0]
}

func GetClientSecret() string {
	clientConfig := getClientConfig()
	return clientConfig[1]
}
