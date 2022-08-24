package backend

import (
	"bytes"
	"github.com/DarkMiMolle/TechnicalTest_Owlint/util"
	"net/http"
)

/*
Client is a representation of a Service (API) we may ask something // NOTE: rename it Service ?
*/
type Client struct {
	// Url is the API's one
	Url string

	// RetryPolicy is the Policy to apply on request error.
	RetryPolicy RetryPolicy
}

// Post allows to make an http.Post request with the RetryPolicy of the Client. Its returns are the InComing returns of the http.Post
func (c Client) Post(contentType string, body []byte) (util.InComing[*http.Response], util.InComing[error]) {
	return c.RetryPolicy.RunPolicy(func() (*http.Response, error) {
		return http.Post(c.Url, contentType, bytes.NewBuffer(body))
	})
}
