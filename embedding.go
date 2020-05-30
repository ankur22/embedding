package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cenkalti/backoff/v4"
)

type retryableHTTPClient struct {
	*http.Client
	b backoff.BackOff
}

func isTransientError(c int) bool {
	if c >= 500 {
		return true
	}
	return false
}

func (r *retryableHTTPClient) Get(url string) (resp *http.Response, err error) {
	err = backoff.Retry(func() error {
		resp, err = r.Client.Get(url)
		if err != nil {
			// Don't retry
			return nil
		}
		if isTransientError(resp.StatusCode) {
			// Let's retry the call
			return fmt.Errorf("Temporary error: status %d", resp.StatusCode)
		}
		// Success, don't retry
		return nil
	}, r.b)
	return
}

func main() {
	c := retryableHTTPClient{http.DefaultClient, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 10)}
	if resp, err := c.Get("https://google.com"); err != nil {
		log.Fatalf("Unexpected error: %v", err)
	} else {
		log.Printf("Successful call to https://google.com: status %d\n", resp.StatusCode)
	}
}
