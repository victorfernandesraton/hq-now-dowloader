package extract

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ErrorOnRequest        = "HTTP request with error"
	ErrorRequestExecution = "HTTP request with error on execute"
	ErrorOnParseJSON      = "JSON response with error"
	ErrorMarshallJSON     = "JSON response with error on marshall"
)

func Request(url string, payload *strings.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorOnRequest, err.Error()))
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorRequestExecution, err.Error()))
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v", ErrorOnParseJSON, err.Error()))
	}

	return body, nil
}
