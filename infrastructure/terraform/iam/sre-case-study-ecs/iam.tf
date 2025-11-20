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

data "aws_iam_policy_document" "ecs_task_assume_role" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
    actions = ["sts:AssumeRole"]
  }
}

data "aws_iam_policy_document" "ecs_task_role" {
  statement {
    effect = "Allow"
    actions = [
      "dynamodb:PutItem",
      "dynamodb:GetItem",
      "dynamodb:UpdateItem",
      "dynamodb:Scan",
    ]
    resources = [
      "arn:aws:dynamodb:${var.region}:${data.aws_caller_identity.current.account_id}:table/${data.terraform_remote_state.dynamodb.outputs.dynamodb_table_name}"
    ]
  }
}

resource "aws_iam_policy" "ecs_task_policy" {
  name   = "ScsEcsPolicy"
  policy = data.aws_iam_policy_document.ecs_task_role.json
}

resource "aws_iam_role" "ecs_task" {
  name               = "ScsEcsRole"
  assume_role_policy = data.aws_iam_policy_document.ecs_task_assume_role.json
}

resource "aws_iam_role_policy_attachment" "ecs_task_attachment" {
  role       = aws_iam_role.ecs_task.name
  policy_arn = aws_iam_policy.ecs_task_policy.arn

}
