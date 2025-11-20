terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "alb/terraform.tfstate"
    region = "eu-west-1"
  }
}