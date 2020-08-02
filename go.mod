module github.com/gphotosuploader/google-photos-api-client-go

go 1.12

replace github.com/gphotosuploader/googlemirror => ../go-googlephotos

require (
	github.com/gphotosuploader/googlemirror v1.0.0
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.19.0
)
