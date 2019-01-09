package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Method   string
	URL      string
	Body     interface{}
	Header   map[string]string
	Username string
	Password string
}

func request(req *Request, res interface{}) error {
	body := bytes.NewBuffer(nil)
	if req.Body != nil {
		data, err := json.Marshal(req.Body)
		if err != nil {
			return err
		}

		body.Write(data)
	}

	request, err := http.NewRequest(req.Method, req.URL, body)
	if err != nil {
		return err
	}

	request.SetBasicAuth(req.Username, req.Password)
	for k, v := range req.Header {
		request.Header.Add(k, v)
	}

	c := http.Client{}
	response, err := c.Do(request)
	if err != nil {
		return err
	}

	if res != nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, res); err != nil {
			return err
		}
	}

	if err := response.Body.Close(); err != nil {
		return err
	}

	if response.StatusCode >= http.StatusMultipleChoices {
		return errors.New("failed")
	}

	return nil
}
