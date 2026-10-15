# Queue Lab

A scaffold for learning Kubernetes, Terraform and Helm through a background
word-count service. Infrastructure is provided; job processing remains TODO.

## Layout

| Directory | Purpose |
| --- | --- |
| cmd/api, cmd/worker | Separate executable entry points |
| internal | Health server, job model, queue/processor placeholders |
| charts/queue-lab | Helm Deployments, Services, ConfigMap, Redis StatefulSet/PVC |
| infra/kind | Local two-node cluster configuration |
| infra/terraform/local | Namespace + Helm release ownership |
| exercises/manifests | Space for your first plain YAML exercise |
| docs | Architecture decisions and learning milestones |
| examples | Mock job submission payload |

## Prerequisites

Go >=1.24, Docker CLI with a running Linux container runtime, kind, kubectl,
Helm 3 and Terraform >=1.6,<2. On macOS Docker can use Colima; verify docker info.
Use a current compatible kind release and its default Kubernetes node image.
No cloud account or registry is needed. Commands below run from this directory.

## Start with application mocks

```sh
make run-api
# In another terminal:
make run-worker
curl -i localhost:8080/healthz
curl -i -H 'Content-Type: application/json' --data @examples/create-job.json localhost:8080/jobs
```

Health endpoints return 200. Job endpoints return 501 until implemented.
Worker health runs on 8081 locally; in Kubernetes each Pod uses 8080.
Environment variables are listed in .env.example; Go does not load that file automatically.

## Deploy locally

Run each step in order and inspect the output:

```sh
make cluster-up
make image-build
make image-load
make helm-check
make init
make plan
make apply
make status
make port-forward
```

Then curl localhost:8080/healthz. Terraform targets kind-queue-lab explicitly,
regardless of your active kubectl context. If using a custom kubeconfig, set
kubeconfig_path in infra/terraform/local/terraform.tfvars.
Terraform apply uses the saved plan; rerun make plan after configuration changes.
After changing Go code, rebuild and load the image, then restart the API/worker
Deployments or use a new image tag. Prefer new tags for repeatable deployments.

## Inspect and experiment

```sh
kubectl --context kind-queue-lab -n queue-lab-dev logs deployment/queue-lab-api
kubectl --context kind-queue-lab -n queue-lab-dev describe pod queue-lab-redis-0
kubectl --context kind-queue-lab -n queue-lab-dev get configmap queue-lab-config -o yaml
helm template queue-lab charts/queue-lab -f charts/queue-lab/values-staging.yaml
```

Change worker.replicas in values-dev.yaml, then make plan and make apply.
See docs/learning-plan.md for the implementation exercises.

## Teardown

```sh
make destroy
make cluster-down
```

Destroy Terraform resources before deleting the cluster. Both teardown steps
can delete lab data; Redis persistence is for Pod restarts, not backups.

## Verification status

This scaffold was structurally checked during creation. Go, Helm, Terraform,
Docker and Kubernetes runtime checks require the tools above and have not been
executed in the creation environment. Run make check and make helm-check first,
then terraform validate after make init. No end-to-end processing is implemented.

## References

- https://kind.sigs.k8s.io/docs/user/quick-start/
- https://helm.sh/docs/chart_template_guide/
- https://registry.terraform.io/providers/hashicorp/helm/latest/docs
- https://redis.io/docs/latest/develop/data-types/streams/
