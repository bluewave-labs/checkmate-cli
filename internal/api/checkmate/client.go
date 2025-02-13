package checkmate

import (
	"net/http"

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
