package venstar

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultTimeout = 2 * time.Second

const (
	apiinfo       = ""
	queryinfo     = "query/info"
	querysensors  = "query/sensors"
	queryruntimes = "query/runtimes"
	control       = "control"
	settings      = "settings"
)

type VenstarClient struct {
	client   *http.Client
	base_url string
	ip       string
	port     int
}

func NewVenstarClient(host string, port int) *VenstarClient {
	c := &VenstarClient{
		ip:       host,
		port:     port,
		client:   &http.Client{Timeout: defaultTimeout},
		base_url: fmt.Sprintf("http://%s:%d", host, port),
	}

	return c
}

func (c *VenstarClient) get(url string, rep VenstarReply) (err error) {
	// TODO: maybe wrap error in non-generic error?
	var resp *http.Response

	if resp, err = c.client.Get(url); err != nil {
		return err
	}

	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		// TODO: return non-generic error
		return fmt.Errorf("http request returned status %s", resp.Status)
	}

	if err := json.Unmarshal(body, rep); err != nil {
		return err
	}

	return nil

}

func (c *VenstarClient) makeUrl(ep string) string {
	return fmt.Sprintf("%s/%s", c.base_url, ep)
}

// func (c *VenstarClient) post(url string) (body []byte, status int, err error) {
// 	return nil, 0, nil
// }

func (c *VenstarClient) Info() (rep InfoReply, err error) {
	err = c.get(c.makeUrl(queryinfo), &rep)
	return
}

func (c *VenstarClient) Sensors() (rep SensorsReply, err error) {
	err = c.get(c.makeUrl(querysensors), &rep)
	return
}
