package checkmate

import (
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
	credentials *config.Credentials
	httpClient  *http.Client
}

func (c *CheckmateClient) SendRequest(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

func NewCheckmateClient(credentials *config.Credentials, httpClient *http.Client) *CheckmateClient {
	return &CheckmateClient{
		credentials: credentials,
		httpClient:  httpClient,
	}
}

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
	default:
		if response.StatusCode >= 400 && response.StatusCode < 500 {
			return types.ErrClientError
		} else if response.StatusCode >= 500 {
			return types.ErrServerError
		}
	}

	return nil
}

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
