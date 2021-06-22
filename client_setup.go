package file_v1

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
)



func NewClient(ctx context.Context, clientId string, clientSecret string, tokenUrl string) *APIClient {

	cfg := &clientcredentials.Config{
		ClientID:    clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenUrl,
	}

	cfg.Client(ctx)

	configuration := NewConfiguration()
	configuration.HTTPClient = cfg.Client(ctx)

	return NewAPIClient(configuration)
}
