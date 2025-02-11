//Package govalent provides the binding for Covalent Rest APIs.
package govalent

import (
	"github.com/cshields143/govalent/class_a"
	"github.com/cshields143/govalent/class_b"
	"github.com/cshields143/govalent/client"
	"net/http"
)

//constant used for API client
const (
	APIURL             = "https://api.covalenthq.com/v1/"
)

//Default HTTP client
var httpClient = &http.Client{}

// APIKey is Covalent API Key.
var APIKey string

// ClassA uses endpoint without client.
func ClassA() *class_a.Client {
	api := client.New(APIURL, APIKey, httpClient)
	return &class_a.Client{API: *api}
}

// ClassB uses endpoint without client.
func ClassB() *class_b.Client {
	api := client.New(APIURL, APIKey, httpClient)
	return &class_b.Client{API: *api}
}

//Client is the Covalent client. It contains all resources available.
type Client struct {
	ClassA class_a.Client
	ClassB class_b.Client
}

//Init initializes the covalent client with given API key, secret key.
func (c *Client) Init(apiKey string) {
	api := client.New(APIURL, apiKey, httpClient)
	c.ClassA = class_a.Client{API: *api}
	c.ClassB = class_b.Client{API: *api}
}
