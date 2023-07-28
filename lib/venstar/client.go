package venstar

import (
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

func (c *VenstarClient) get(url string) (body []byte, status int, err error) {
	var resp *http.Response
	resp, err = c.client.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if body, err = io.ReadAll(resp.Body); err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, err
}

func (c *VenstarClient) makeUrl(ep string) string {
	return fmt.Sprintf("%s/%s", c.base_url, ep)
}

// func (c *VenstarClient) post(url string) (body []byte, status int, err error) {
// 	return nil, 0, nil
// }

func (c *VenstarClient) Info() (rep InfoReply, err error) {
	body, status, err := c.get(c.makeUrl(queryinfo))
	if err != nil {
		return rep, err
	}
	if status != http.StatusOK {
		return rep, fmt.Errorf("Http request returned status %d", status)
	}
}

func (c *VenstarClient) Sensors() SensorsReply {
	return SensorsReply{}
}
