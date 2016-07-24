package payload

import (
	"io/ioutil"
	"net/http"
)

// A Collection provides request object.
type Collection struct {
	Version   string     `json:"version"`
	Token     string     `json:"token"`
	Instances []Instance `json:"instances"`
}

// An Instance is.
type Instance struct {
	URL string `json:"url"`
}

// Get gets bytes of URL over request.
func (d Instance) Get() ([]byte, error) {
	client := &http.Client{}

	resp, err := client.Get(d.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
