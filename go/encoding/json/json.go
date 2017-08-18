package json

import (
	"bytes"
	"encoding/json"
)

var meta = map[string]interface{}{
	"version": "v1.0.1",
}

func merge_v1(v interface{}) interface{} {
	var data = make(map[string]interface{})
	data["meta"] = meta

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	return data
}

func merge_v2(v interface{}) interface{} {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	data["meta"] = meta
	return data
}

func merge_v3(v interface{}) interface{} {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	err = json.NewDecoder(&buf).Decode(&data)
	if err != nil {
		return err
	}

	data["meta"] = meta
	return data
}
