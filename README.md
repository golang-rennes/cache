Various implementations of a **cache** mechanism.

Code samples and slides of the presentation of March 14th, 2018, by Benoît Masson in the [Golang Rennes](https://www.meetup.com/Golang-Rennes/events/248219655/) context.
These samples are given for demonstrating and exercising, and should not be used as such in production.

# Organization

The cache interface to implement is defined as follows in `src/cache/cache.go`:

```golang
type T interface {
	Get(key string) ([]byte, bool)
	Add(key string, content []byte)
	Invalidate(key string)
}
```

This interface is implemented by 6 structures, in the corresponding file:

1. `cache.none`: dummy implementation, does nothing;
1. `cache.nemory`: in-memory implementation based on Go’s maps;
1. `cache.syncMemory`: thread-safe version of the latter;
1. `cache.file`: persistent on-disk cache;
1. `cache.expirable`: in-memory cache, items are removed after their lifetime expires;
1. `cache.bounded`: bounded-size in-memory cache.

# Execution

#### Tests

Basic tests for all cache implementations are defined in `cache/cache_test.go`, and can be run using:

```sh
go test ./...
```

#### Web server

These implementations can also be tested individually using the given demonstration web server (folder `server`).
This server replies with the reversed request path, with a 1 second delay if the result does not exist in cache, immediately otherwise.

Run it on default port 8888 with:

```sh
go install server && bin/server
```

and query it in another terminal with `curl`:

```sh
curl http://localhost:8888/test && echo
curl http://localhost:8888/test/path && echo
curl http://localhost:8888/test/third/path && echo
```

The server uses the `cache.New` function to initialize its cache, change the alias in file `cache/cache.go` to use a different cache implementation.
