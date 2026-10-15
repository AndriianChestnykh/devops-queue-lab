# Design decisions

## Scope
A word-count job queue with POST /jobs and GET /jobs/{id}. No frontend.
The API and worker are separate Go binaries in one module and container image.
Both expose health endpoints; business endpoints deliberately return HTTP 501.
The worker does not consume jobs yet. Readiness initially means HTTP availability.

## Technology choices
- Go 1.24 language baseline, standard net/http and slog: minimal application dependencies.
  This is a reproducible baseline, not a claim about the latest Go release.
- Planned Redis adapter: go-redis/v9, added when queue logic is implemented.
- Redis Streams consumer group: at-least-once processing, acknowledgements,
  pending entries and crash recovery. XAUTOCLAIM and idempotent result writes
  are required before claiming reliable job recovery.
- Redis also stores job status/results initially; no separate SQL database.
- Single Redis StatefulSet with PVC and AOF everysec. This tolerates Pod replacement,
  not node or cluster loss. A sudden crash can lose roughly a second of AOF writes.
  This is a learning deployment with no Redis authentication or network isolation.
- kind: two nodes (control plane + worker) inside the local container runtime.
  Docker CLI/runtime required; on macOS Colima is sufficient.
- Helm: one small first-party chart, no chart dependencies.
- Terraform: stable Kubernetes provider 2.x and Helm provider 3.x syntax,
  local state initially. Commit .terraform.lock.hcl after terraform init.

## Ownership
kind CLI creates/deletes the cluster. Terraform owns the namespace and Helm release.
Helm owns Deployments, Services, ConfigMap and Redis StatefulSet.
Do not use helm upgrade or kubectl apply for permanent changes to these resources.
Use kubectl for observation and explicit failure experiments, then restore configuration.

## Planned processing flow
1. API validates text, generates ID and stores a queued record, then XADDs the job.
2. Worker XREADGROUPs, counts words and stores the result before XACK.
3. A recovery loop claims stale pending entries; processing by job ID is idempotent.
4. GET reads status and result. Define retention limits for streams and result keys.

Atomic enqueue + status creation, consumer identity, retry limits and poison messages
are deliberate future exercises. No durability or delivery guarantee is implemented yet.

## Storage lifecycle
StatefulSet PVCs normally outlive release removal. Terraform namespace deletion removes
PVCs in that namespace; deleting kind destroys its local volumes. Treat both teardown
paths as deletion of lab data. PVC retention is not a backup.

## Deferred features
Secrets/authentication, Ingress/TLS, NetworkPolicy, HPA/metrics-server, remote Terraform
state and cloud cluster provisioning. Add these as separate learning milestones.
