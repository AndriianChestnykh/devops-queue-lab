# Learning milestones

1. Build/run the mock API and worker; inspect the Dockerfile.
2. Create kind and implement the plain manifests in exercises/manifests.
   Inspect labels/selectors, endpoints, DNS, probes, logs and PVC binding.
3. Read and render the chart. Compare dev vs staging with helm template.
4. Deploy with Terraform. Read the plan and inspect Terraform state ownership.
5. Implement job submission, result lookup and word counting with Redis Streams.
6. Kill a worker after delivery but before acknowledgement; implement pending recovery.
7. Scale workers to three via values-dev.yaml and Terraform; prove each job completes.
8. Delete the Redis Pod and check persisted results. Then rehearse clean teardown.
9. Add Secret-based Redis auth, NetworkPolicy and API Ingress.
10. Add autoscaling, remote state and finally a Terraform-managed cloud cluster.

Acceptance for the scaffold: health endpoints respond, /jobs returns 501,
API and worker Pods become ready, Redis PVC is bound. End-to-end jobs are a later milestone.
