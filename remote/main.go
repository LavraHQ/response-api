package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lavrahq/response-api/config"
)

func initHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
}

func initDataServiceRequest(method string, path string, body io.Reader) *http.Request {
	fullURL := fmt.Sprintf("%s/%s", config.Config.DataURL, path)
	req, _ := http.NewRequest(method, fullURL, body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-Hasura-Admin-Secret", config.Config.DataSecret)

	return req
}

// QueryRequestArgs is the QueryRequest arguments for a request.
type QueryRequestArgs struct {
	Type string                 `json:"type"`
	Args map[string]interface{} `json:"args"`
}

// NewQueryRequest initializes a new http.Request for a Data Service
// metdata query request.
func NewQueryRequest(args *QueryRequestArgs) ([]byte, error) {
	client := initHTTPClient()
	data, _ := json.Marshal(args)
	req := initDataServiceRequest("POST", "v1/query", bytes.NewBuffer(data))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
