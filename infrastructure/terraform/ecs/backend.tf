terraform {
  backend "s3" {
    bucket = "yahlatci-terraform-states"
    key    = "ecs/terraform.tfstate"
    region = "eu-west-1"
  }
}