CLUSTER := queue-lab
CONTEXT := kind-queue-lab
IMAGE := queue-lab:dev
TF_DIR := infra/terraform/local

.PHONY: help fmt check run-api run-worker cluster-up image-build image-load helm-check init plan apply status port-forward destroy cluster-down
help:
	@echo "Start: make cluster-up image-build image-load helm-check init plan apply"
fmt:
	gofmt -w cmd internal
	terraform fmt -recursive infra/terraform
check:
	go vet ./...
	go test ./...
run-api:
	go run ./cmd/api
run-worker:
	HTTP_ADDR=:8081 go run ./cmd/worker
cluster-up:
	kind create cluster --name $(CLUSTER) --config infra/kind/cluster.yaml --wait 120s
image-build:
	docker build -t $(IMAGE) .
image-load:
	kind load docker-image $(IMAGE) --name $(CLUSTER)
helm-check:
	helm lint charts/queue-lab -f charts/queue-lab/values-dev.yaml
	helm template queue-lab charts/queue-lab -f charts/queue-lab/values-dev.yaml > /tmp/queue-lab-rendered.yaml
init:
	terraform -chdir=$(TF_DIR) init
plan:
	terraform -chdir=$(TF_DIR) plan -out=local.tfplan
apply:
	terraform -chdir=$(TF_DIR) apply local.tfplan
status:
	kubectl --context $(CONTEXT) -n queue-lab-dev get pods,svc,pvc
port-forward:
	kubectl --context $(CONTEXT) -n queue-lab-dev port-forward svc/queue-lab-api 8080:8080
destroy:
	terraform -chdir=$(TF_DIR) destroy
cluster-down:
	kind delete cluster --name $(CLUSTER)
