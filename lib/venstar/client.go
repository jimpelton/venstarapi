package venstar

import (
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeout = 2 * time.Second

var client = &http.Client{Timeout: defaultTimeout}

func Client() *http.Client {
	return client
}

func Get(url string) (body []byte, status int, err error) {
	var resp *http.Response
	resp, err = client.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, err
}
