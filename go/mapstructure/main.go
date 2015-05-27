package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
)

type BodyA struct {
	AccessToken string `json:"access_token" mapstructure:"access_token"`
	Campaign    string `json:"campaign"`
}

type BodyB struct {
	Access_Token string `json:"access_token"`
	Campaign     string `json:"campaign"`
}

func main() {
	spew.Dump(process(&BodyA{}))
	spew.Dump(process(&BodyB{}))
}

func process(b interface{}) interface{} {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           b,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}
	query := map[string]interface{}{
		"access_token": "test_token",
		"campaign":     "test_campaign",
	}
	if err = decoder.Decode(query); err != nil {
		panic(err)
	}
	return b
}
