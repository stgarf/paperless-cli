package paperless

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// HTTPClient is an HTTP client to interact with the Paperless API server
var HTTPClient http.Client

func (p Paperless) String() string {
	if !p.UseHTTPS {
		return fmt.Sprintf("http://%v:%v%v", p.Hostname, p.Port, p.Root)
	}
	return fmt.Sprintf("https://%v:%v%v", p.Hostname, p.Port, p.Root)
}

func (p Paperless) showInstanceInformation() string {
	return fmt.Sprintf("Username: %v, Hostname: %v, Port: %v, API root: %v, HTTPS: %v", p.Username, p.Hostname, p.Port, p.Root, p.UseHTTPS)
}

// ShowInstanceInformation shows the currently loaded Paperless instance configuration
func (p Paperless) ShowInstanceInformation() {
	fmt.Println(p.showInstanceInformation())
}

// ReturnHTTPResponse takes an http.Response pointer and returns a ([]bytes, error)
func ReturnHTTPResponse(r *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err == nil {
		return b, nil
	}
	log.Errorf("%v", err)
	return []byte{}, err
}

// Authenticate to a Paperless API instance
func (p Paperless) Authenticate() *http.Request {
	p.Request, _ = http.NewRequest("", "", nil)
	p.Request.Header.Set("User-Agent", "paperless-cli")
	p.Request.SetBasicAuth(p.Username, p.Password)
	return p.Request
}

func (p Paperless) switchResponse(r *http.Response) ([]byte, error) {
	switch r.StatusCode {
	case 200:
		resp, err := ReturnHTTPResponse(r)
		if err != nil {
			return []byte{}, err
		}
		return resp, nil
	case 403:
		req := p.Authenticate()
		req.Method = "GET"
		url := &url.URL{}
		url, _ = url.Parse(p.String())
		req.URL = url
		r, err := HTTPClient.Do(req)
		if err != nil {
			return []byte{}, err
		}
		resp, err := ReturnHTTPResponse(r)
		if err != nil {
			return []byte{}, err
		}
		return resp, nil
	default:
		fmt.Println("HTTP Status Code:", r.StatusCode)
		resp, err := ReturnHTTPResponse(r)
		if err != nil {
			return []byte{}, err
		}
		return resp, nil
	}
}

// MakeRequest .
func (p Paperless) MakeRequest(method string) ([]byte, error) {
	switch method {
	case "GET":
		resp, err := HTTPClient.Get(p.String())
		if err != nil {
			return []byte{}, err
		}
		data, err := p.switchResponse(resp)
		if err != nil {
			return []byte{}, err
		}
		return data, nil
	case "POST":
		resp, err := HTTPClient.Post(p.String(), "application/json", nil)
		if err != nil {
			return []byte{}, err
		}
		data, err := p.switchResponse(resp)
		if err != nil {
			return []byte{}, err
		}
		return data, nil
	default:
		resp, err := HTTPClient.Get(p.String())
		if err != nil {
			return []byte{}, err
		}
		data, err := p.switchResponse(resp)
		if err != nil {
			return []byte{}, err
		}
		return data, nil
	}
}
