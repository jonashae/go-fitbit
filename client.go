package fitbit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/jonashae/go-fitbit/internal/api"
	"github.com/jonashae/go-fitbit/internal/auth"
)

const baseURL = "https://api.fitbit.com"
const timeout = 30 * time.Second

// Client is used to access the Fitbit API
type Client struct {
	authHandler *auth.Handler
}

// NewClient creates new API client
func NewClient(config Config) *Client {
	tokenStorage := auth.NewFileTokenStorage(config.TokenPath)

	authHandler := auth.NewHandler(
		config.ClientID,
		config.ClientSecret,
		config.Scopes,
		config.CallbackServerPort,
		tokenStorage,
	)

	return &Client{authHandler: authHandler}
}



func (client Client) getJson(path string) gjson.Result {
	context, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()



	httpClient, err := client.authHandler.Client(context)
	if err != nil {
	}


	url := client.createURL(path, nil)

	fmt.Println(url)

	response, err := httpClient.Get(url.String())
	if err != nil {

	}

	var json gjson.Result

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json = gjson.Parse(string(bodyBytes))
		return json
	}

	defer response.Body.Close()
	return json
}



func (client Client) getResource(path string, params url.Values, out interface{}) error {
context, cancel := context.WithTimeout(context.Background(), timeout)
defer cancel()

httpClient, err := client.authHandler.Client(context)
if err != nil {
return err
}

url := client.createURL(path, params)

response, err := httpClient.Get(url.String())
if err != nil {
return err
}

defer response.Body.Close()
return client.handleResponse(response, out)
}

func (client Client) handleResponse(response *http.Response, out interface{}) error {
	statusCode := response.StatusCode
	decoder := json.NewDecoder(response.Body)

	if statusCode == http.StatusOK {
		return decoder.Decode(out)
	}

	var result api.GenericResponse
	decoder.Decode(&result)

	if len(result.Errors) == 0 {
		log.Printf("received status code %d\n", statusCode)
	}

	for _, err := range result.Errors {
		log.Printf("[%d] %s: %s\n", statusCode, err.Type, err.Message)
	}

	if statusCode == http.StatusUnauthorized {
		client.authHandler.Reset()
		return errors.New("authorization error")
	}

	if statusCode == http.StatusTooManyRequests {
		return errors.New("too many requests")
	}

	if statusCode >= 400 && statusCode < 500 {
		return errors.New("bad request")
	}

	return errors.New("unknown error")
}

func (client Client) createURL(path string, params url.Values) *url.URL {
	url, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	url.Path = path + ".json"
	url.RawQuery = params.Encode()

	return url
}
