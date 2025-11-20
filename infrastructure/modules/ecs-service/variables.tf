
variable "region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "eu-west-1"
}

variable "image_uri" {
  description = "The ECR image URI for the ECS service"
  type        = string
}

variable "image_tag" {
  description = "The tag of the Docker image for the ECS service"
  type        = string
}

variable "target_group_arn" {
  description = "The ARN of the target group to attach the ECS service to"
  type        = string
}

variable "app_env" {
  description = "The application environment (e.g., local, dev, prod)"
  type        = string
  default     = "local"
}

variable "task_role_arn" {
  description = "The ARN of the IAM role for the ECS task"
  type        = string
}

variable "name" {
  description = "The base name for resources"
  type        = string
  default     = "sre-case-study"
}

variable "ecs_cluster_arn" {
  description = "The ARN of the ECS cluster"
  type        = string
}

variable "alb_security_group_id" {
  description = "The security group ID of the ALB"
  type        = string
  
}

variable "vpc_private_subnet_ids" {
  description = "List of private subnet IDs in the VPC"
  type        = list(string)
  
}