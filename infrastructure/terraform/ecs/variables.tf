
variable "region" {
  description = "The AWS region to deploy resources in"
  type        = string
  default     = "eu-west-1"

}


variable "dynamodb_table_name" {
  description = "The name of the DynamoDB table"
  type        = string
  default     = "picusv3"

}