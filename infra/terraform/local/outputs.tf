output "namespace" {
  value = kubernetes_namespace_v1.lab.metadata[0].name
}

output "release_name" {
  value = helm_release.lab.name
}
