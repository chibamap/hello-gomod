package fuga

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)


func Welcome() (string, error) {

	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(ip) == 0 {
		return "", ErrNoIP
	}

	return fmt.Sprintf("welcome, %v", string(ip)), nil
}

