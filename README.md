# Maker Monitor

This project aims to make useful Maker and Dai system metrics easily available through open source monitoring tools.
I have chosen to use Prometheus as the backend for simplicity, and Grafana for the frontend for flexibility.

## Developing

This project uses Docker, Kubernetes, and [Tilt](https://tilt.dev/) to try and get local development as close to the
deployment as possible. Please make sure you have these tools installed before moving forward. Once you're setup run:

```
# This is secrect, and isn't shared in the repo, which is why you need to create it yourself
k create configmap monitor-config --from-literal=INFURA_PROJECT_ID=<your infura project ID>

# This stores the prometheus config into a k8s config map
k create configmap prometheus-config --from-file=prometheus/prometheus.yml

# Start the stack!
tilt up
```

### Building Monitor

```
./monitor/scripts/build.sh
```

*Note:* This creates a Docker images tagged `ddaws/maker-monitor:latest`

## To Do

- Update Prometheus deployment to mount data volume
- Add Grafana (of course!)
- Update Docker image to build more efficiently (I think it's go get'ing on build)

## Deployment

To deploy Maker Monitor you'll need to create a series of config maps, pods and services via Kuberenetes. The following
set of commands assumes you've aliased `kubectl` as `k`.

```bash
# This is secrect, and isn't shared in the repo, which is why you need to create it yourself
k create configmap monitor-config --from-literal=INFURA_PROJECT_ID=<your infura project ID>

# Deploy the monitor service
k apply -f monitor/k8s/pod.yml
k apply -f monitor/k8s/service.yml

# Deploy the prometheus service
k create configmap prometheus-config --from-file=prometheus/prometheus.yml
k apply -f prometheus/k8s/pod.yml
```

If you didn't run into any errors you should be able to connect to Prometheus via

```bash
k port-forward prometheus-pod 9090:9090
```

## Useful Links

- [Go Ethereum Book](https://goethereumbook.org/en/)
- [Go Prometheus Client Lib](https://godoc.org/github.com/prometheus/client_golang/prometheus)
- [Annotated Contracts via Hypothesis](https://via.hypothes.is/https://github.com/makerdao/dss/blob/master/src/vat.sol)