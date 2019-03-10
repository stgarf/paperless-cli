package paperless

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
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
func (p Paperless) MakeGetRequest(urlString string) ([]gjson.Result, error) {
	log.Debugf("GET: %v", urlString)

	// Create a client and authenticated request
	client := http.Client{Timeout: time.Second * 5}
	nextURL := urlString
	results := []gjson.Result{}

	for nextURL != "" {
		log.Debugln("Gettings results...")
		req := ReturnAuthenticatedRequest(p.Username, p.Password)
		req.Method = "GET"
		urlPtr, _ := url.Parse(nextURL)
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
			return []gjson.Result{}, errors.New(s)
		}

		json := gjson.ParseBytes(b)
		nextURL = json.Get("next").String()
		moreResults := json.Get("results").Array()
		if len(moreResults) == 0 {
			results = json.Get("name").Array()
		}
		for _, res := range moreResults {
			results = append(results, res)
		}

	}
	return results, nil
}

// GetNameByID returns a name for a given ID path
// e.g. "/api/correspondents/4/"
func (p Paperless) GetNameByID(path string) string {
	p.Root = path
	u := fmt.Sprint(p)
	results, err := p.MakeGetRequest(u)
	if err != nil {
		log.Errorf("An error occurred making request: %v", err.Error())
	}
	var s string
	for _, result := range results {
		s += result.String()
	}
	return s
}
