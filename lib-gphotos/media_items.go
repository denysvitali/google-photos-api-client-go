package gphotos

import (
	"context"
	"github.com/denysvitali/go-googlephotos/api/photoslibrary/v1"
)


func (c *Client) MediaItemsByAlbum(ctx context.Context, album *photoslibrary.Album, limit int) (items []*photoslibrary.MediaItem, err error) {
	mediaItemsCall := c.MediaItems.Search(&photoslibrary.SearchMediaItemsRequest{
		AlbumId: album.Id,
		PageSize: defaultPageSize})
	smr, err := mediaItemsCall.Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	// Fetch all pages
	mediaItems := make([]*photoslibrary.MediaItem, 0)
	currentMediaItems := smr.MediaItems

	var nextPageToken = smr.NextPageToken

	for ;; {
		if len(currentMediaItems) == 0 || smr == nil {
			break
		}

		if limit != 0 {
			if len(mediaItems)+len(currentMediaItems) > limit {
				break
			}
		}

		mediaItems = append(mediaItems, currentMediaItems...)

		if nextPageToken == "" {
			break
		}

		c.log.Debugf("getting next page... (fetched %d elements)", len(mediaItems))

		mediaItemsCall = c.MediaItems.Search(&photoslibrary.SearchMediaItemsRequest{
			AlbumId: album.Id,
			PageSize: defaultPageSize,
			PageToken: nextPageToken})

		smr, err = mediaItemsCall.Context(ctx).Do()
		if err != nil {
			c.log.Error(err)
			return nil, err
		}

		nextPageToken = smr.NextPageToken
		currentMediaItems = smr.MediaItems
	}

	return mediaItems, nil
}