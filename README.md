# racey-go

Mini HTTP server, written to accompany the blog post [Getting started with Goroutines and channels](https://jldec.me/getting-started-with-go-part-3-goroutines-and-channels). The server listens on port 3000, counts requests, and returns the count with each request.

### To run

```sh
go run .
```

### To run with race detection

```sh
export GORACE=halt_on_error=1
go run -race .
```

### To trigger the race detection

This uses xargs to spawn parallel curl requests.

```
cat urls.txt | xargs -P 4 -n 1 curl
```

### To test with node.js

A similar Node.js implementation is provided for comparison.

```
node server.js
```