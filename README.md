# Quick n' Dirty Go HTTP Server

Simple Go HTTP server that offers a couple of routes to test/benchmark systems using HTTP GET/POST methods.

## Build and Run

- run `go build`
- run `./go-simple-http-api` (will listen in port 3000)


## TODO

- Evaluate replacing atomics with channels
- Cleanup code to use packages to clean up
- Accept some options using https://github.com/spf13/cobra
- Provide endpoint for stats (for now /api/now offers total requests)
- Add TLS support

### Performance GET `/api/now`

Command Used: 

```shell
bombardier localhost:3000/api/now -c 125 -l
```

```shell
Statistics        Avg      Stdev        Max
  Reqs/sec    345765.68   21585.15  389791.16
  Latency      360.52us   250.94us    66.22ms
  Latency Distribution
     50%   190.00us
     75%   417.00us
     90%     0.89ms
     95%     1.23ms
     99%     2.16ms
  HTTP codes:
    1xx - 0, 2xx - 3453839, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    91.56MB/s
```

### Performance POST `/api/jsonpayload` (8KB payload)

Command Used:

```shell
bombardier -c 150 -m POST -H 'Content-Type: application/json' -b '{"ID":1,"Data":"<generate 8KB payload somehow>"}' http://localhost:3000/api/jsonpayload
```

```shell
Statistics        Avg      Stdev        Max
  Reqs/sec     80536.77    4432.21   87613.76
  Latency        1.86ms     1.71ms    76.42ms
  HTTP codes:
    1xx - 0, 2xx - 803901, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   649.08MB/s
```
