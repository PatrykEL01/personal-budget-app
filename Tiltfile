load('ext://restart_process', 'docker_build_with_restart')

## build the docker image and restart the container when the source code changes

docker_build_with_restart(
    'patrykel1/personal-budget',
    '.',
    dockerfile='Dockerfile',
    entrypoint='/personal-budget',
    live_update=[
        sync('./', '/app'),
        run('cd /app && go build -o /personal-budget'),
    ]
)


k8s_yaml('k8s/postgresql/operator-manifest.yaml')
k8s_yaml('k8s/postgresql/secret.yaml')
k8s_yaml('k8s/postgresql/postgres-instance.yaml')


# Load the Kubernetes yaml files

k8s_yaml('k8s/app/namespace.yaml')
k8s_yaml('k8s/app/service.yaml')
k8s_yaml('k8s/app/deployment.yaml')
k8s_resource('personal-budget-app', port_forwards=8080)


allow_k8s_contexts('kind-local-k8s')




