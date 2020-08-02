package google_photos_api_client_go

import (
	"context"
	"fmt"
	photoslibrary "github.com/denysvitali/go-googlephotos"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/api/googleapi"
)

func (c *Client) retryableMediaItemBatchCreateDo(ctx context.Context, request *photoslibrary.BatchCreateMediaItemsRequest, filename string) (*photoslibrary.BatchCreateMediaItemsResponse, error) {
	var res *photoslibrary.BatchCreateMediaItemsResponse
	var err error

	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		c.log.Debugf("Sending media item creation: file=%s, retry=%d", filename, i)

		res, err = c.MediaItems.BatchCreate(request).Context(ctx).Do()
		if err == nil {
			// If there is not an error, it doesn't need to be retried.
			return res, nil
		}

		// handle retries
		if e, ok := err.(*googleapi.Error); ok {
			switch {
			case e.Code == http.StatusTooManyRequests:
				// Rate limit error. Minimum 60s delay.
				after, err := strconv.ParseInt(e.Header.Get("Retry-After"), 10, 64)
				if err != nil || after == 0 {
					after = 60
				}

				c.log.Infof("Media creation. Rate limit reached, sleeping for %d seconds: file=%s", after, filename)

				time.Sleep(time.Duration(after) * time.Second)
				continue
			case e.Code >= http.StatusInternalServerError && e.Code <= http.StatusNetworkAuthenticationRequired:
				// Retryable 500 error.
				// TODO: It should be exponential backoff
				c.log.Errorf("Media creation. Received error, sleeping for 10 seconds before retrying: file=%s", filename)

				time.Sleep(10 * time.Second)
				continue
			}
		}
		return nil, fmt.Errorf("unexpected error response: file=%s, err=%s", filename, err)

	}
	return res, fmt.Errorf("too many retries: file=%s, err=%s", filename, err)
}
