output "ecr_repository_url" {
  description = "URL of the ECR repository."
  value       = aws_ecr_repository.my_app.repository_url
}

output "ecs_cluster_name" {
  description = "Name of the ECS cluster."
  value       = aws_ecs_cluster.my_cluster.name
}

# Add any other outputs you need

