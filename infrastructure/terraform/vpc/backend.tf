terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "vpc/terraform.tfstate"
    region = "eu-west-1"
  }
}