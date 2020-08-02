package gphotos_test

import (
	gphotos "github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos"
	"github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos/internal/log"
	"github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos/internal/uploader"
	"net/http"
	"testing"
)

type mockUploadSessionStore struct{}

func (m *mockUploadSessionStore) Get(f string) []byte {
	return []byte(f)
}

func (m *mockUploadSessionStore) Set(f string, u []byte) {}

func (m *mockUploadSessionStore) Delete(f string) {}

func TestNewClientWithResumableUploads(t *testing.T) {
	c := http.DefaultClient
	store := &mockUploadSessionStore{}

	t.Run("EmptyHTTPClient", func(t *testing.T) {
		_, err := gphotos.NewClientWithOptions(nil, store)
		if err == nil {
			t.Errorf("NewClientWithOptions error was expected here")
		}
	})

	t.Run("WithNilUploadSessionStore", func(t *testing.T) {
		_, err := gphotos.NewClientWithOptions(c, nil)
		if err != uploader.ErrNilStore {
			t.Errorf("NewClientWithOptions - error was expected here: got=%s, want=%s", err, uploader.ErrNilStore)
		}
	})

	t.Run("WithoutOptions", func(t *testing.T) {
		got, err := gphotos.NewClientWithOptions(c, store)
		if err != nil {
			t.Errorf("NewClientWithOptions - error was not expected here: err=%s", err)
		}
		if got.Service == nil {
			t.Errorf("NewClientWithOptions - Photos service was not created")
		}
	})

	t.Run("WithOptionLog", func(t *testing.T) {
		l := log.NewDiscardLogger()
		got, err := gphotos.NewClientWithOptions(c, store, gphotos.WithLogger(l))
		if err != nil {
			t.Errorf("NewClientWithOptions - error was not expected here: err=%s", err)
		}
		if got.Service == nil {
			t.Errorf("NewClientWithOptions - Photos service was not created")
		}
	})
}