data "terraform_remote_state" "lambda_iam" {
    backend = "s3"
    
    config = {
        bucket = "yahlatci-terraform-states"
        key    = "iam/sre-case-study-lambda/terraform.tfstate"
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
module "lambda" {
    source = "../../modules/lambda"
    lambda_role_arn = data.terraform_remote_state.lambda_iam.outputs.lambda_role_arn
    target_group_arn = data.terraform_remote_state.alb.outputs.lambda_target_group_arn
    app_env = "prod"
    image_uri = "887495603804.dkr.ecr.eu-west-1.amazonaws.com/sre-case-study-lambda"
    image_tag = var.image_tag
}