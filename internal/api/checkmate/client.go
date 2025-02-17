package checkmate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/bluewave-labs/checkmate-cli/internal/api/checkmate/types"
	"github.com/bluewave-labs/checkmate-cli/internal/config"
)

// CheckmateClient is a client for the Checkmate API.
// It is used to send HTTP requests to the Checkmate API.
// See [Checkmate Server OpenAPI Specs] for more details.
//
// [Checkmate Server OpenAPI Specs]: https://github.com/bluewave-labs/Checkmate/blob/develop/Server/openapi.json
type CheckmateClient struct {
	credentials   *config.Credentials
	httpClient    *http.Client
	authenticator *Authenticator
}

func (c *CheckmateClient) performRequest(method string, path string, payload any, auth *Authenticator) (*types.APIResponse, error) {
	var requestBody []byte
	var err error
	var req *http.Request

	if method == http.MethodGet {
		req, err = http.NewRequest(method, c.credentials.APIBaseURL+path, nil)
	} else {
		requestBody, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body := bytes.NewBuffer(requestBody)
		req, err = http.NewRequest(method, c.credentials.APIBaseURL+path, body)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(auth.headerKey, auth.header+c.credentials.APIKey)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if err := responseStatusCodeHandler(response); err != nil {
		return nil, err
	}

	return bodyParser(response.Body)
}

func NewCheckmateClient(credentials *config.Credentials, httpClient *http.Client, authenticator *Authenticator) *CheckmateClient {
	return &CheckmateClient{
		credentials:   credentials,
		httpClient:    httpClient,
		authenticator: authenticator,
	}
}

// Checkmate Client has pre-defined error objects for common HTTP status codes.
// It returns an error object based on the HTTP status code of the response.
// If no error is found, it returns nil.
func responseStatusCodeHandler(response *http.Response) error {
	if response == nil {
		return errors.New("response is nil")
	}

	switch response.StatusCode {
	case http.StatusForbidden:
		return types.ErrForbidden
	case http.StatusUnauthorized:
		return types.ErrUnauthorized
	case http.StatusNotFound:
		return types.ErrNotFound
	case http.StatusBadRequest:
		return types.ErrBadRequest
	case http.StatusUnprocessableEntity:
		return types.ErrUnprocessableEntity
	}

	return nil
}

// Checkmate API returns a structured response in JSON format.
//
// Successful response:
//
//	success: bool
//	msg: string
//	data: any (nullable)
//
// Error response:
//
//	success: bool
//	msg: string
func bodyParser(body io.ReadCloser) (*types.APIResponse, error) {
	if body == nil {
		return nil, errors.New("response body is nil")
	}
	defer body.Close()

	var apiResponse types.APIResponse
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
