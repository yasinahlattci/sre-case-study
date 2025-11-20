
variable "region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "eu-west-1"

}

variable "name" {
  description = "The base name for resources"
  type        = string
  default     = "sre-case-study"
  
}