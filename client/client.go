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
	ErrorIncorrectBodyFormat = errors.New("incorrect body format")
	ErrorSendRequest         = errors.New("couldn't send request")
	ErrorUnknown             = errors.New("unknown error")
)

type HTTPStatusError struct {
	status int
}

func NewHTTPStatusError(status int) error {
	return &HTTPStatusError{status}
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("status code: %d ", e.status)
}

func (e *HTTPStatusError) Status() int {
	return e.status
}

func PostJson(client *http.Client, url string, body string) (err error) {
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
