variable "aws_region" {
  description = "AWS region where resources will be created."
  default     = "us-east-1" # Change to your desired region
}

variable "aws_account_id" {
  description = "Your AWS account ID."
  type        = string
}

variable "my_app_image" {
  description = "Docker image for my app."
  type        = string
}

# Add any other variables you need

