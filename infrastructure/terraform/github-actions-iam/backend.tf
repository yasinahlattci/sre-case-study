terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "github-actions-iam/terraform.tfstate"
    region = "eu-west-1"
  }
}