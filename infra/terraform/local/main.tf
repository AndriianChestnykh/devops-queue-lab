resource "kubernetes_namespace_v1" "lab" {
  metadata {
    name = var.namespace
  }
}

resource "helm_release" "lab" {
  name      = "queue-lab"
  namespace = kubernetes_namespace_v1.lab.metadata[0].name
  chart     = "${path.module}/../../../charts/queue-lab"
  values    = [file("${path.module}/../../../charts/queue-lab/values-dev.yaml")]
  wait      = true
  atomic    = true
  timeout   = 300
}
