terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "sre-case-study/terraform.tfstate"
    region = "eu-west-1"
  }
}