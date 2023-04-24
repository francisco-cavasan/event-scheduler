package Services

import (
	"errors"
	"net/http"
	"net/url"
)

func Handle(content string, address string) error {
	// Create the request body
	postBody := url.Values{}
	postBody.Add("address", address)

	// Create the POST request
	request, err := http.PostForm("http://localhost:8081/send", postBody)
	if err != nil {
		return err
	}

	defer request.Body.Close()

	// Check the response status code
	if request.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + request.Status)
	}

	return nil
}
