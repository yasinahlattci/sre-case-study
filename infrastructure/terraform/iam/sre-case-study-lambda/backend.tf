terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "iam/sre-case-study-lambda/terraform.tfstate"
    region = "eu-west-1"
  }
}