package file_v1

import (
	"context"
	apic "github.com/antinvestor/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const ctxKeyService = "filesClientKey"

func defaultFilesClientOptions() []apic.ClientOption {
	return []apic.ClientOption{
		apic.WithEndpoint("https://files.api.antinvestor.com"),
	}
}

func ToContext(ctx context.Context, fileClient *FilesClient) context.Context {
	return context.WithValue(ctx, ctxKeyService, fileClient)
}

func FromContext(ctx context.Context) *FilesClient {
	fileClient, ok := ctx.Value(ctxKeyService).(*FilesClient)
	if !ok {
		return nil
	}

	return fileClient
}

// FilesClient is a client for interacting with the file service API.
// Methods, except Close, may be called concurrently. However,
//fields must not be modified concurrently with method calls.
type FilesClient struct {
	APIClient
	// The x-ant-* metadata to be sent with each request.
	xMetadata metadata.MD
}

// NewFilesClient creates a new file client.
// The service that an application uses to create and read files stored in the system
func NewFilesClient(ctx context.Context, opts ...apic.ClientOption) (*FilesClient, error) {
	clientOpts := defaultFilesClientOptions()

	httClient, err := apic.HttpClient(ctx, append(clientOpts, opts...)...)

	if err != nil {
		return nil, err
	}

	configuration := NewConfiguration()
	configuration.HTTPClient = httClient

	cl := &FilesClient{}
	cl.cfg = configuration
	cl.common.client = &cl.APIClient
	cl.DefaultApi = (*DefaultApiService)(&cl.common)

	cl.setClientInfo()

	return cl, nil
}

// setClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (fc *FilesClient) setClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", apic.VersionGo()}, keyval...)
	kv = append(kv, "grpc", grpc.Version)
	fc.xMetadata = metadata.Pairs("x-ai-api-client", apic.XAntHeader(kv...))
}



