module "lambda" {
  source         = "terraform-aws-modules/lambda/aws"
  version        = "8.1.2"
  function_name  = "${var.name}-lambda"
  create_package = false
  publish        = true
  image_uri = "${var.image_uri}:${var.image_tag}"
  package_type = "Image"
  create_role = false
  lambda_role = var.lambda_role_arn
  memory_size = 128

  allowed_triggers = {
    alb = {
      principal  = "elasticloadbalancing.amazonaws.com"
      source_arn = var.target_group_arn
    }
  }
  environment_variables = {
    APP_ENV = var.app_env
  }
}

resource "aws_alb_target_group_attachment" "lambda_attachment" {
  target_group_arn = var.target_group_arn
  target_id        = module.lambda.lambda_function_arn
}