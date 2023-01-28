package netlas

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// Client struct for making requests to the netlas API
type Client struct {
	APIKey string
	Client *retryablehttp.Client
}

type DomainCount struct {
	Count int `json:"count"`
}

type Domain struct {
	Name string `json:"domain"`
}

type Domains struct {
	Items []struct {
		Data Domain `json:"data"`
	} `json:"items"`
}

// NewClient creates a new netlas client with the provided API key
func NewClient(apiKey string) *Client {
	client := retryablehttp.NewClient()
	client.RetryMax = 2
	client.RetryWaitMin = 1 * time.Second
	client.RetryWaitMax = 5 * time.Second
	client.Logger = nil

	return &Client{
		APIKey: apiKey,
		Client: client,
	}
}

// GetDomainCount makes a request to the netlas API to get the total domain count for a given query
func (c *Client) GetDomainCount(domain string) (*DomainCount, error) {
	req, err := retryablehttp.NewRequest("GET", "https://app.netlas.io/api/domains_count/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("domain:*.%s", domain))
	q.Add("indices", "")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("x-api-key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse response and return domain count
	var dc DomainCount
	if err := json.NewDecoder(resp.Body).Decode(&dc); err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	return &dc, nil
}

// GetSubDomains makes a request to the netlas API to get a list of domains for a given query and starting index
func (c *Client) GetSubdomains(domain string, start int) (*Domains, error) {
	req, err := retryablehttp.NewRequest("GET", "https://app.netlas.io/api/domains/", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("domain:*.%s", domain))
	q.Add("start", strconv.Itoa(start))
	q.Add("indices", "")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("x-api-key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// parse response and return list of domains
	var d Domains
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, fmt.Errorf("error decoding response: %s", err)
	}

	return &d, nil
}

func GetAllSubdomains(c *Client, domain string, outputCh *chan string) error {
	// Get the total domain count for the given domain
	dc, err := c.GetDomainCount(domain)
	if err != nil {
		return err
	}

	// Iterate through the total domain count
	for i := 0; i < dc.Count; i += 20 {
		// Get a list of subdomains for the current index
		d, err := c.GetSubdomains(domain, i)
		if err != nil {
			return err
		}

		// Add the returned subdomains to the list
		for _, subdomain := range d.Items {
			*outputCh <- subdomain.Data.Name
		}
	}
	close(*outputCh)

	return nil
}
