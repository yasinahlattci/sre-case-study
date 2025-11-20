
data "terraform_remote_state" "vpc" {
    backend = "s3"
    config = {
        bucket = "yahlatci-terraform-states"
        key    = "vpc/terraform.tfstate"
        region = "eu-west-1"
    }
}

data "terraform_remote_state" "iam" {
    backend = "s3"
    config = {
        bucket = "yahlatci-terraform-states"
        key    = "iam/sre-case-study-ecs/terraform.tfstate"
        region = "eu-west-1"
    }
}

    data "terraform_remote_state" "ecs" {
        backend = "s3"
        config = {
            bucket = "yahlatci-terraform-states"
            key    = "ecs/terraform.tfstate"
            region = "eu-west-1"
        }
    }

data "terraform_remote_state" "alb" {
    backend = "s3"
    config = {
        bucket = "yahlatci-terraform-states"
        key    = "alb/terraform.tfstate"
        region = "eu-west-1"
    }
}

module "ecs-service" {
    source = "../../modules/ecs-service"
    app_env = "prod"
    vpc_private_subnet_ids = data.terraform_remote_state.vpc.outputs.private_subnets
    task_role_arn = data.terraform_remote_state.iam.outputs.ecs_task_role_arn
    ecs_cluster_arn = data.terraform_remote_state.ecs.outputs.ecs_cluster_arn
    target_group_arn = data.terraform_remote_state.alb.outputs.api_target_group_arn
    alb_security_group_id = data.terraform_remote_state.alb.outputs.alb_security_group_id
    image_uri = "887495603804.dkr.ecr.eu-west-1.amazonaws.com/sre-case-study-api:b3c1f401fdfe6a199590c662014295c26057a11e"
}