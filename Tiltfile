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
# Apply PostgreSQL CRDS (cloudnative-pg)
local("kubectl apply -f https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/release-1.20/releases/cnpg-1.20.6.yaml")

# Apply PostgreSQL manifests
k8s_yaml(['k8s/postgresql/secret.yaml', 'k8s/postgresql/postgres-instance.yaml'])

# Load the Kubernetes global yaml files

k8s_yaml('k8s/global/namespace.yaml')

# Apply the app yaml files
k8s_yaml(['k8s/app/deployment.yaml', 'k8s/app/service.yaml'])




## Allow the Tiltfile to be run in a kind cluster
allow_k8s_contexts('kind-local-k8s')


k8s_resource(
    'personal-budget-app',
    port_forwards=8080
)


k8s_resource(
    new_name='DB_SECRET',
    objects=['personal-budget-user-secret']
)

k8s_resource(
    new_name='postgresql',
    objects=['personal-budget-db-cluster']
)

k8s_resource(
    new_name='namespace',
    objects=["personal-budget-app:namespace"]
)
