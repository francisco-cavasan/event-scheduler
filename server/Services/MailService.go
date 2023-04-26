package Services

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PostBody struct {
	Address string `json:"address"`
	PetName string `json:"petName"`
}

func Handle(petName string, address string) error {
	fmt.Println("send email to", address)
	if address == "" {
		return errors.New("address is required")
	}
	// Create the request body
	postBody := []byte(`{"address":"` + address + `", "petName":"` + petName + `"}`)

	// Create the POST request
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8081/send", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(request)
	// Check the response status code
	if err != nil && res == nil {
		return errors.New("unexpected status code: " + res.Status)
	}

	return nil
}
