terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "ecr/terraform.tfstate"
    region = "eu-west-1"
  }
}