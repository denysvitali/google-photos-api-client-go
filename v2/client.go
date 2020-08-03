package google_photos_api_client_go

import (
	"net/http"
	"sync"

	"golang.org/x/oauth2"

	photoslibrary "github.com/denysvitali/go-googlephotos"

	"github.com/denysvitali/google-photos-api-client-go/v2/internal/log"
	"github.com/denysvitali/google-photos-api-client-go/v2/internal/uploader"
)

// Client is a client for uploading a media.
// photoslibrary does not provide `/v1/uploads` API so we implement here.
type Client struct {
	// Google Photos client
	*photoslibrary.Service
	// Uploader to upload new files to Google Photos
	uploader *uploader.Uploader

	log log.Logger
	mu  sync.Mutex

	token *oauth2.Token // DEPRECATED: `token` will disappear in the next MAJOR version.
}

// NewClientWithOptions constructs a new gphotos.Client from the provided HTTP client and
// the given options.
//
// `httpClient` is an client with authentication credentials.
// `store` is an UploadSessionStore to keep upload sessions to resume uploads.
func NewClientWithOptions(httpClient *http.Client, store UploadSessionStore, options ...Option) (*Client, error) {
	photosService, err := photoslibrary.New(httpClient)
	if err != nil {
		return nil, err
	}

	upldr, err := uploader.NewUploader(httpClient, uploader.WithResumableUploads(store))
	if err != nil {
		return nil, err
	}

	c := &Client{
		Service:  photosService,
		uploader: upldr,
		log:      log.NewDiscardLogger(),
	}

	for _, opt := range options {
		opt(c)
	}

	return c, nil
}

// WithLogger set a new Logger to log messages.
func WithLogger(l log.Logger) func(*Client) {
	return func(c *Client) {
		c.log = l
	}
}

// Option defines an option for a Client
type Option func(*Client)
type UploadSessionStore uploader.UploadSessionStore

type MemoryUploadSessionStore struct {
	store map[string][]byte
}
var _ UploadSessionStore = (*MemoryUploadSessionStore) (nil)

func (m MemoryUploadSessionStore) Get(fingerprint string) []byte {
	return m.store[fingerprint]
}

func (m MemoryUploadSessionStore) Set(fingerprint string, url []byte) {
	m.store[fingerprint] = url
}

func (m MemoryUploadSessionStore) Delete(fingerprint string) {
	delete(m.store, fingerprint)
}