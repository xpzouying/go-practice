# map vs sync/map

Golang 1.9 introduce new map in multiple thread safe.

I write two map examples for multiple thread safe, use `map` and `sync/Map`.

- map/main.go
    > Use sync/mutex to ensure multiple threads safe
- sync_map/main.go
    > multiple threads safe insides
