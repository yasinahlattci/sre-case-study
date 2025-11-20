terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "dynamodb/terraform.tfstate"
    region = "eu-west-1"
  }
}