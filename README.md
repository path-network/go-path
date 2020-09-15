# Go Path

![alt text](img/path.png "Path logo")


The official client library written in Go to interact with Path.net's API

## Installation

The package can be downloaded using Go's builtin get command:

`go get github.com/path-network/go-path/path`

After installing go-path, you may use it in your project by importing it

```go
import "github.com/path-network/go-path/path"
```

## Getting Started
To start using Go Path, you must first request an access token. You may either authenticate using account credentials, or a client ID and client secret:

```go
// Request a new OAuth2 token with our credentials
tokenReq := path.AccessTokenRequest{
	Username: "foo",
	Password: "bar",
}

// Instantiate a new Path API client using our credentials. An access token will be obtained and used in subsequent
// requests to other endpoints
client, err := path.NewClient(tokenReq)

if err != nil {
	log.Fatalf("Error authenticating: %s\n", err.Error())
}
```

After successful authentication, you may access endpoints that require authorization:

```go
rules, err := path.GetRules()

if err != nil {
	log.Fatalf("Error fetching rules: %s\n", err.Error())
}

. . .

```

## Documentation
For reference on how to use this package, please refer to the [documentation](https://godoc.org/github.com/path-network/go-path/path).
