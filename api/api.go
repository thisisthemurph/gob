package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// DatabricksAPI is a generic API for making calls to Databricks
type DatabricksAPI struct {
	BaseURL string
	token   string
}

// NewDatabricksAPI creates a new DatabricksAPI
func NewDatabricksAPI(baseURL, token string) DatabricksAPI {
	return DatabricksAPI{
		BaseURL: baseURL,
		token:   token,
	}
}

// Get sends a GET request to the Databricks API at the given endpoint
// and returns a []bytes obtect containing the response.
//
// The method handles setting the Authorization Bearer token.
//
// Example:
//
//	databricks := NewDatabricksAPI(BaseURL, Token)
//	b, err := databricks.Get("/list")
func (a DatabricksAPI) Get(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, a.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status error")
	}

	return io.ReadAll(resp.Body)
}
