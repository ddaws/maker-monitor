# Maker Monitor

This project aims to make useful Maker and Dai system metrics easily available through open source monitoring tools.
I have chosen to use Prometheus as the backend for simplicity, and Grafana for the frontend for flexibility.

## Developing

### Installing Dependancies

Just run it! We're using Go modules :D 

### Building the Monitor binary

```
./monitor/scripts/build.sh
```

*Note:* The binary build script builds for the `golang:alpine` Docker container. Use `go build` or `go run` for local.

## To Do

- Add Grafana (of course!)

- Update Go Dockerfile to use multistage builds
https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

## Useful Links

- [Go Ethereum Book](https://goethereumbook.org/en/)
- [Go Prometheus Client Lib](https://godoc.org/github.com/prometheus/client_golang/prometheus)
- [Annotated Contracts via Hypothesis](https://via.hypothes.is/https://github.com/makerdao/dss/blob/master/src/vat.sol)