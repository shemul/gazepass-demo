package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)
type AuthUser struct {
	ID    *uuid.UUID `json:"id,omitempty"`
	Email string     `json:"email,omitempty"`
}

type AuthResponse struct {
	User     *AuthUser `json:"user"`
	Verified bool      `json:"verified"`
}

var (
	API_KEY = "9866de97-2c85-49ea-9b66-6ddf83d0e6a6"
	SECRET_KEY = "2b716d02-fd93-59c6-b4c1-a99a0947b141"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/pvt" , func(writer http.ResponseWriter, request *http.Request) {
		AuthToken := request.Header.Get("Authorization")
		//hostname, _ := os.Hostname()
		println(AuthToken)
		resp, err := VerifyAccessToken(AuthToken, API_KEY, SECRET_KEY)
		if err != nil {
			println(err.Error())
			writer.WriteHeader(http.StatusForbidden)
		}else {
			marshal, err := json.Marshal(resp.User)
			if err != nil {
				println(err.Error())
			}
			writer.Write(marshal)
		}
	})

	http.ListenAndServe("localhost:3000" ,nil)
	println("server running on 3000")
}


func VerifyAccessToken(access_token, api_key, api_secret string) (*AuthResponse, error) {
	reqDataJSON, err := json.Marshal(map[string]string{
		"access_token": access_token,
		"api_key":      api_key,
		"api_secret":   api_secret,
	})

	if err != nil {
		return nil, errors.New("Couldn't create JSON")
	}

	resp, err := http.Post("https://api.gazepass.com/user/auth", "application/json", bytes.NewBuffer(reqDataJSON))
	if err != nil {
		return nil, errors.New("Couldn't make HTTP request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Error from Gazepass API %v" ,resp.StatusCode))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Couldn't read body")
	}

	var authResponse AuthResponse
	err = json.Unmarshal(bodyBytes, &authResponse)
	if err != nil {
		return nil, errors.New("Couldn't parse JSON")
	}

	return &authResponse, nil
}