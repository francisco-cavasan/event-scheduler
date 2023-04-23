package Services

import (
	"bytes"
	"errors"
	"net/http"
)

func Handle(content string, address string) error {
	// Create a new HTTP client
	client := &http.Client{}

	// Create the request body
	requestBody := []byte(`{"content":"` + content + `", "address":"` + address + `"}`)

	// Create the POST request
	request, err := http.NewRequest("POST", "http://localhost:8081", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Send the POST request
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + response.Status)
	}

	// Handle successful response
	// ...

	return nil
}
