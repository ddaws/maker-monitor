# Checkout Tiltfiles! https://tilt.dev/

docker_build('ddaws/maker-monitor', './monitor/', dockerfile='monitor/Dockerfile')

k8s_yaml([
    'monitor/k8s/pod.yml',
    # 'monitor/k8s/service.yml',
    # 'prometheus/k8s/pod.yml',
])

k8s_resource('monitor-pod', port_forwards=8080)
# k8s_resource('prometheus-pod', port_forwards=9090)
