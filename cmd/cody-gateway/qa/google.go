package qa

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sourcegraph/sourcegraph/internal/codygateway"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

type GoogleGatewayFeatureClient struct{}

func (o GoogleGatewayFeatureClient) GetRequest(f codygateway.Feature, req *http.Request, stream bool) (*http.Request, error) {
	if f == codygateway.FeatureCodeCompletions {
		return nil, errNotImplemented
	}
	if f == codygateway.FeatureChatCompletions {
		body := fmt.Sprintf(`{"model":"gemini-1.5-flash-latest","contents":[{"parts":[{"text":"You are Cody"}],"role":"user"},{"parts":[{"text":"Ok I am Cody"}],"role":"model"},{"parts":[{"text":"Write Bubble sort in Python"}],"role":"user"}],"stream":%t}`, stream)
		req.Method = "POST"
		req.URL.Path = "/v1/completions/google"
		req.Body = io.NopCloser(strings.NewReader(body))
		return req, nil
	}
	return nil, errors.New("unknown feature: " + string(f))
}
