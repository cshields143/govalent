# Govalent

govalent is an awesome go client for Covalent Rest APIs, originally built by @zaebee

But it has two issues:

- 80 second timeout on HTTP calls; this is eliminated--if I wait, I wait
- API keys are included in the URL instead of using Basic Auth; I guess this isn't really an issue, but I feel better knowing my key is buried in a header instead of in the URL itself

For further documentation (these two changes affect nothing the user sees or interacts with), see https://github.com/zaebee/govalent

## Changes that DO affect what the user sees or interacts with

- Added the Block Heights endpoint from Class A
- A general `PaginateParams` type has been added, and `LogEventsParams` changed to allow specifying page size
- Don't log our GET requests

## Install

```sh
go get github.com/cshields143/govalent
```
