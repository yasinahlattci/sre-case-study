terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "iam/github-actions/terraform.tfstate"
    region = "eu-west-1"
  }
}