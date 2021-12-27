package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorIncorrectJson(t *testing.T) {
	err := PostJson(http.DefaultClient, "http://google.com", `{{"data": 1}`)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorIncorrectBodyFormat)
}

func TestErrorSendRequest(t *testing.T) {
	err := PostJson(http.DefaultClient, "http://incorrect.url", `{}`)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, ErrorSendRequest)
}

func TestIncorrectStatus(t *testing.T) {
	err := PostJson(http.DefaultClient, "http://httpstat.us/500", `{}`)

	assert.NotNil(t, err)
	e, ok := err.(*HTTPStatusError)
	assert.Equal(t, ok, true)
	assert.Equal(t, e.Status(), http.StatusInternalServerError)

}

func TestOk(t *testing.T) {
	err := PostJson(http.DefaultClient, "http://httpstat.us/200", `{}`)
	assert.Nil(t, err)
}
