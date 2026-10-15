variable "kubeconfig_path" {
  type    = string
  default = "~/.kube/config"
}

variable "kube_context" {
  type    = string
  default = "kind-queue-lab"
}

variable "namespace" {
  type    = string
  default = "queue-lab-dev"
}
