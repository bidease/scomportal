package scomportal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (api *API) request(method string, path string, out interface{}, body interface{}) error {
	var req *http.Request
	var err error

	if body != nil {
		var bytesData []byte
		bytesData, err = json.Marshal(body)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, fmt.Sprintf("%s/%s", api.BaseURL, path), bytes.NewBuffer(bytesData))
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%s/%s", api.BaseURL, path), nil)
		if err != nil {
			return err
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-Email", api.Email)
	req.Header.Set("X-User-Token", api.Token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.Unmarshal(respBody, &out)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) getRequest(partURL string, out interface{}) error {
	return api.request("GET", partURL, &out, nil)
}
