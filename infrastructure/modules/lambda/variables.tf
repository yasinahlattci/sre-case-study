
variable "region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "eu-west-1"
}

variable "lambda_image" {
  description = "The ECR image URI for the Lambda function"
  type        = string
}

variable "target_group_arn" {
  description = "The ARN of the target group to attach the Lambda function to"
  type        = string
}

variable "app_env" {
  description = "The application environment (e.g., local, dev, prod)"
  type        = string
  default     = "local"
}

variable "lambda_role_arn" {
  description = "The ARN of the IAM role for the Lambda function"
  type        = string
}

variable "name" {
  description = "The base name for resources"
  type        = string
  default     = "sre-case-study"
}