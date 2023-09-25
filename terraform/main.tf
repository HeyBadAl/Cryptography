provider "aws" {
  region = "us-east-1" # Change to your desired region
}

resource "aws_ecr_repository" "my_app" {
  name = "my-encryption-registry"
}

resource "aws_ecs_cluster" "my_cluster" {
  name = "my-encryption-cluster"
}

resource "aws_ecs_task_definition" "my_app" {
  family                   = "my-encryption-registry"
  network_mode             = "bridge"
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn

  container_definitions = jsonencode([{
    name  = "my-encryption-container"
    image = "${aws_account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/my-encryption-registry:latest"
    portMappings = [{
      containerPort = 8080
      hostPort      = 8080
    }]
  }])
}

resource "aws_iam_role" "ecs_execution_role" {
  name = "ecs_execution_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

data "aws_region" "current" {}

