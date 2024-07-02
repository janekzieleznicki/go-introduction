package generator

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log/slog"
	"io/ioutil"
)

type Password struct {
	Password string
}

func GetPassword(len uint32) string {
	requestURL := fmt.Sprintf("https://api.genratr.com/?length=%d&uppercase&lowercase", 31+len)
	// fmt.Printf("URL: %s\n", requestURL)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		slog.Error("client: could not create request: %s\n", err)
		return ""
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("client: error making http request: %s\n", err)
		return ""
	}


	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		slog.Error("client: could not read response body: %s\n", err)
		return ""
	}

	var pass Password
	err = json.Unmarshal(resBody, &pass)
	if err != nil {
		slog.Error("client: could not unmarshall json %s\n", err)
		return ""
	}
	slog.Debug("Got password from server: ", "pass",pass)

	return pass.Password
}

func GeneratePasswords() chan string {
	channel := make(chan string)
	go func() {
		var count uint32 = 0
		for {
			channel <- GetPassword(count)
			count = count +1
		}
		slog.Info("Generator destroyed")
	}()
	return channel
}
