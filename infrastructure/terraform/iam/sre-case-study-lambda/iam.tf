data "aws_caller_identity" "current" {
}

data "terraform_remote_state" "dynamodb" {
  backend = "s3"

  config = {
    bucket = "yahlatci-terraform-states"
    key    = "dynamodb/terraform.tfstate"
    region = "eu-west-1"
  }
}

data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}



data "aws_iam_policy_document" "lambda_role" {
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:DeleteItem",
      "dynamodb:GetItem",
    ]
    resources = [
      "arn:aws:dynamodb:${var.region}:${data.aws_caller_identity.current.account_id}:table/${data.terraform_remote_state.dynamodb.outputs.dynamodb_table_name}"
    ]
  }
}


resource "aws_iam_policy" "lambda_policy" {
  name   = "ScsLambdaPolicy"
  policy = data.aws_iam_policy_document.lambda_role.json
}

resource "aws_iam_role" "lambda_role" {
  name               = "ScsLambdaRole"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

resource "aws_iam_role_policy_attachment" "lambda_attachment" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}

resource "aws_iam_role_policy_attachment" "lambda_basic_execution" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_ecr_readonly" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
  
}