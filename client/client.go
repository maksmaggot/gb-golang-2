package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

const jsonContentType = "application/json"

var (
	// ErrorIncorrectBodyFormat incorrect body format error
	ErrorIncorrectBodyFormat = errors.New("incorrect body format")
	// ErrorSendRequest send request error
	ErrorSendRequest = errors.New("couldn't send request")
	// ErrorUnknown unknown error
	ErrorUnknown = errors.New("unknown error")
)

// HTTPStatusError http error with status
type HTTPStatusError struct {
	status int
}

// NewHTTPStatusError constructor
func NewHTTPStatusError(status int) error {
	return &HTTPStatusError{status}
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("status code: %d ", e.status)
}

// Status returns http error status
func (e *HTTPStatusError) Status() int {
	return e.status
}

// PostJSON send json to url
func PostJSON(client *http.Client, url string, body string) (err error) {
	var js map[string]interface{}

	defer func() {
		if v := recover(); v != nil {
			err = errors.Wrap(ErrorUnknown, err.Error())
		}
	}()

	err = json.Unmarshal([]byte(body), &js)
	if err != nil {
		return errors.Wrap(ErrorIncorrectBodyFormat, err.Error())
	}

	res, err := client.Post(url, jsonContentType, bytes.NewReader([]byte(body)))
	if err != nil {
		return errors.Wrap(ErrorSendRequest, err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return NewHTTPStatusError(res.StatusCode)
	}

	return nil
}
