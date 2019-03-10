package paperless

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

// Paperless struct represents a Paperless instance
type Paperless struct {
	Hostname string
	UseHTTPS bool
	Port     string
	Root     string
	Username string
	Password string
}

// ReturnAuthenticatedRequest to a Paperless API instance
func ReturnAuthenticatedRequest(u, p string) *http.Request {
	request, _ := http.NewRequest("", "", nil)
	request.Header.Set("User-Agent", "paperless-cli")
	request.SetBasicAuth(u, p)
	return request
}

// MakeGetRequest makes a request of method to url with args.
func MakeGetRequest(creds []string, u string) ([]byte, error) {
	log.Debugf("GET: %v", u)

	// Create a client and authenticated request
	client := http.Client{Timeout: time.Second * 5}
	req := ReturnAuthenticatedRequest(creds[0], creds[1])
	req.Method = "GET"
	urlPtr, _ := url.Parse(u)
	req.URL = urlPtr

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("An error occurred with request: %v", err.Error())
	}

	// Read the response
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Unable to read response body: %v", err.Error())
	}
	if resp.StatusCode != 200 {
		s := fmt.Sprintf("Received non-200 status code: %v: Body: %v", resp.Status, string(b))
		return []byte{}, errors.New(s)
	}

	return b, nil
}
