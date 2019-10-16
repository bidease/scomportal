package scomportal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (api *API) request(method string, path string, out interface{}, body interface{}) *http.Response {
	var req *http.Request
	var err error

	if body != nil {
		var bytesData []byte
		bytesData, err = json.Marshal(body)
		if err != nil {
			log.Fatalln(err)
		}

		req, err = http.NewRequest(method, fmt.Sprintf("%s/%s", api.BaseURL, path), bytes.NewBuffer(bytesData))
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%s/%s", api.BaseURL, path), nil)
		if err != nil {
			log.Fatalln(err)
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-User-Email", api.Email)
	req.Header.Set("X-User-Token", api.Token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	err = json.Unmarshal(respBody, &out)
	if err != nil {
		log.Fatalln(err)
	}

	return res
}

func (api *API) getRequest(partURL string, out interface{}) {
	api.request("GET", partURL, &out, nil)
}
