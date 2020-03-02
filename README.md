# Maker Monitor

This project aims to make useful Maker and Dai system metrics easily available through open source monitoring tools.
I have chosen to use Prometheus as the backend for simplicity, and Grafana for the frontend for flexibility.

## Developing

### Installing Dependancies

```
./monitor/scripts/install.sh
```

*Note:* This package does not use Go modules though will migrate to Go modules in the future

### Building the Monitor binary

```
./monitor/scripts/build.sh
```

*Note:* The binary build script builds for the `golang:alpine` Docker container. Use `go build` or `go run` for local.

## To Do

- Add Grafana (of course!)

- Introduce Go modules to deprecate the use of a simple install script
https://blog.golang.org/using-go-modules

- Update Go Dockerfile to use multistage builds
https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

## Useful Links

- [Go Ethereum Book](https://goethereumbook.org/en/)
