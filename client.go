package gopingdom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const baseUrl = "https://api.pingdom.com/api/3.1"

type JsonError struct {
	Error struct {
		StatusCode   int
		StatusDesc   string
		ErrorMessage string
	}
}

type RestClient struct {
	webClient *http.Client
	token     string
	host      string
	debug     bool
}

func NewRestClient(token string, debug bool, timeoutSeconds time.Duration) (client *RestClient, error error) {
	return &RestClient{
		token: token,
		debug: debug,
		webClient: &http.Client{
			Timeout: time.Second * timeoutSeconds,
		},
	}, error
}

func (r *RestClient) Get(uri string) (data []byte, error error) {

	urlString := baseUrl + uri

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL - %s", urlString)
	}

	//fmt.Println(urlString)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	response, err := r.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RestClient) Delete(uri string) (data []byte, error error) {

	urlString := baseUrl + uri
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL - %s", urlString)
	}

	req, err := http.NewRequest("DELETE", urlString, nil)
	if err != nil {
		return nil, err
	}

	response, err := r.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RestClient) Patch(uri string, json []byte) (data []byte, error error) {

	urlString := baseUrl + uri
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL - %s", urlString)
	}

	payload := bytes.NewReader(json)

	req, err := http.NewRequest("PATCH", urlString, payload)
	if err != nil {
		return nil, err
	}

	response, err := r.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RestClient) Post(uri string, json []byte) (data []byte, error error) {

	urlString := baseUrl + uri
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, fmt.Errorf("Invalid URL - %s", urlString)
	}

	payload := bytes.NewReader(json)

	req, err := http.NewRequest("POST", urlString, payload)
	if err != nil {
		return nil, err
	}

	response, err := r.doRequest(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RestClient) doRequest(req *http.Request) ([]byte, error) {
	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("UserAgent", "gopingdom")
	req.Header.Set("Authorization", "Bearer "+r.token)

	resp, err := r.webClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if r.debug {
		//log.Println("Request: " + req.URL.RequestURI())
		//log.Println("JSON Response:\n" + body)
	}

	if resp.StatusCode > 299 {

		if strings.Contains(string(body), "message") {
			var jec JsonError
			json.Unmarshal(body, &jec)
			return nil, fmt.Errorf("%d - %s - %s", jec.Error.StatusCode, jec.Error.StatusDesc, jec.Error.ErrorMessage)
		}
		return nil, fmt.Errorf("%s", body)

	}

	return body, nil

}
