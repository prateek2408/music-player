package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func getSongList(input string) ([]string, error) {

	result := make([]string, 0)
	addPath := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && Contains(supportedFormats, filepath.Ext(path)) {
			result = append(result, path)
		}
		return nil
	}
	err := filepath.Walk(input, addPath)

	return result, err

}

func getTokenFromFile(input string) (*oauth2.Token, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in the browser, then type the "+
		"authorization code: \n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}
	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrive token from web %v", err)
	}
	return token
}

func getClient(tokenFile string, config *oauth2.Config) *http.Client {
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		//Falling to Read from Web.
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)

}

func saveToken(tokenFile string, token *oauth2.Token) {
	fmt.Printf("Saving Credentials to file %s\n", tokenFile)
	f, err := os.OpenFile(tokenFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to Save Token into File: %v\n", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Contains(arr []string, input string) bool {
	for _, v := range arr {
		if v == input {
			return true
		}
	}
	return false
}
