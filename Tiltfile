# Checkout Tiltfiles! https://tilt.dev/

# File watch and recompile the monitor process for Linux
local_resource(
    'compile-monitor',
    'cd ./monitor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/monitor .',
    deps=[
        './monitor/go.mod',
        './monitor/go.sum',
        './monitor/main.go',
        './monitor/collector',
    ],
)

docker_build('ddaws/maker-monitor', './monitor/', dockerfile='monitor/Dockerfile')

k8s_yaml([
    'monitor/k8s/pod.yml',
    # 'monitor/k8s/service.yml',
    # 'prometheus/k8s/pod.yml',
])

k8s_resource('monitor-pod', port_forwards=8080)
# k8s_resource('prometheus-pod', port_forwards=9090)