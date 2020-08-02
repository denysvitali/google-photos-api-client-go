# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/) and this project adheres to [Semantic Versioning](https://semver.org/).

## 1.1.5
### Changed
- Update required `googlemirror` package version to v0.3.5.

## 1.1.4
> This version was not working due to `photoslibrary` dependency. PLEASE UPDATE TO v1.1.5 ASAP.
### Changed
- Update required `googlemirror` package version to v0.3.4.

## 1.1.3
### Added
- [CONTRIBUTING](CONTRIBUTING.md) guide line has been added.
### Changed
- [README](README.md) has been updated fixing some typos.
- Module use an interface Logger to log activity. This allow to implement different logging systems. See [internal/log/logger.go](v2/internal/log/logger.go).
### Deprecated
- Once Go 1.13 has been published, previous Go 1.11 support is deprecated. This project will maintain compatibility with the last two major versions published.
### Fix
- Fix uploaded photos without a correct file name. ([#32][i32])
### Removed
- Remove progress information when uploading, if you want to have progress bars when upload a file, you should inject a reader to the Uploader().

[i32]: https://github.com/gphotosuploader/google-photos-api-client-go/issues/32

## 1.1.2
### Changed
- Update [golangci](https://github.com/golangci/golangci-lint) linter to version 1.20.0.
### Fixed
- Fix fatal error introduced in the last version. (#28)

## 1.1.1
### Fixed
- Fix race condition on `GetOrCreateAlbumByName()`. Google Photos API allow you to create several albums with the same name. (#26)

## 1.1.0
### Added
- New `NewClientWithResumableUploads()` function to create a Google Photos client with uploads that can be resumed.
- New `AddMediaItem()` method to upload contents and make it available in Google Photos.
- New `uploader` internal package implementing uploads to Google Photos.
### Changed
- Refactor how Google response is checked. It's following Google Photos best practices. (#10)
- Updated README documentation and added examples.
- Refactor how retries are handled by the code. See `retryableMediaItemBatchCreateDo()` method.
### Deprecated
- `NewClient()` function has been deprecated. Use `NewClientWithResumableUploads()` instead.
- `UploadFile()` and `UploadFileResumable()` methods has been deprecated. Use `AddMediaItem()` instead.
- `noserver-gphotos` package has been deprecated.

## 1.0.7
### Added
- A new Logger has been added to be shared across the whole package. See `logger.go`.
- Tests for almost all the code, except `uploads.go` that needs a lot of refactor.
- Package's documentation and examples of use.

### Changed
- `xerrors` is used instead of `errors`. Preparing code for Go 1.13 release.

### Deprecated
- `Token()` method has been deprecated. Current package implementation doesn't need to have OAuth token.

## 1.0.6
### Fixed
- Fix progress logging. (#19)

## 1.0.5
### Added
- Code quality reporting using [Codebeat](https://codebeat.co/projects/github-com-gphotosuploader-google-photos-api-client-go-master).
- Code coverage reporting using [codecov](https://codecov.io/gh/gphotosuploader/google-photos-api-client-go).

### Fixed
- `ReadProgressReporter` was giving panic at some circumstances. There was a problem casting `int64` to `int`. (#17)
- Fix progress calculations for files big sized. 

### Changed
- CI platform is now [drone.io](https://drone.io).

## 1.0.4
### Fixed
- Fix `AlbumByName` to check against all Google Photos album list (#12).

## 1.0.3
### Added
- Add resumable file uploads. You can use new `UploadFileResumable` method to upload files that can be resumed. See [documentation](https://godoc.org/github.com/gphotosuploader/google-photos-api-client-go/lib-gphotos) for more details.

## 1.0.2
### Added
- Add makefile for easy test and linting
- Add CI using travis-ci.com
- Add travis and goodocs badges to README

## 1.0.1
### Added
- Add semantic versioning to this package
- Add Go module support
- Add MIT license

### Changed
- Clean up of useless files / directories
- Update package documentation

### Fixed
- Fix issue #8 on parent repository [here](https://github.com/nmrshll/google-photos-api-client-go/issues/8)

### Removed
- Removed `Makefile` support

## 1.0.0
Initial release after clone it from [original repository](https://github.com/nmrshll/google-photos-api-client-go). Latest commit was [3dac07f](https://github.com/nmrshll/google-photos-api-client-go/commit/3dac07f1b07f249ac4a9805b9d60afe0f68c34b2)
