module "vpc" {
    source  = "terraform-aws-modules/vpc/aws"
    version = "v6.5.1"

    name = "sre-case-study-vpc"
    azs = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
    cidr = "10.0.0.0/16"
    private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
    public_subnets  = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]
    enable_flow_log = false
    single_nat_gateway = true
    enable_nat_gateway = true
}