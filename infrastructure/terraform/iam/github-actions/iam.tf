
resource "aws_iam_openid_connect_provider" "github" {
  url = "https://token.actions.githubusercontent.com"

  client_id_list = [
    "sts.amazonaws.com"
  ]

  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
}

data "aws_iam_policy_document" "github_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Federated"
      identifiers = [aws_iam_openid_connect_provider.github.arn]
    }

    actions = ["sts:AssumeRoleWithWebIdentity"]

    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }

    condition {
      test = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values = [
        "repo:${var.github_username}/${var.github_repo}:*",
      ]
    }
  }
}

resource "aws_iam_role" "github_actions" {
  name               = "GithubActionsOidcRole"
  assume_role_policy = data.aws_iam_policy_document.github_assume_role.json

}

data "aws_iam_policy_document" "github_permissions" {
  statement {
    sid    = "ECR"
    effect = "Allow"
    actions = [
      "ecr:GetDownloadUrlForLayer",
      "ecr:BatchGetImage",
      "ecr:BatchCheckLayerAvailability",
      "ecr:DescribeRepositories",
      "ecr:GetAuthorizationToken",
      "ecr:CreateRepository",
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:PutImage",
      "ecr:BatchDeleteImage",
      "ecr:SetRepositoryPolicy",
    ]
    resources = ["*"]
  }

  statement {
    sid    = "ECS"
    effect = "Allow"
    actions = [
      "ecs:CreateService",
      "ecs:UpdateService",
      "ecs:DeleteService",
      "ecs:DescribeServices",
      "ecs:RegisterTaskDefinition",
      "ecs:DeregisterTaskDefinition",
      "ecs:DescribeTaskDefinition",
      "ecs:RunTask",
      "ecs:StopTask",
      "ecs:ListTasks",
      "ecs:DescribeTasks",
      "ecs:ListClusters",
      "ecs:DescribeClusters"
    ]
    resources = ["*"]
  }

  statement {
    sid    = "Lambda"
    effect = "Allow"
    actions = [
      "lambda:CreateFunction",
      "lambda:UpdateFunctionCode",
      "lambda:UpdateFunctionConfiguration",
      "lambda:DeleteFunction",
      "lambda:PublishVersion",
      "lambda:GetFunction",
      "lambda:AddPermission",
      "lambda:RemovePermission",
      "lambda:InvokeFunction",
      "lambda:ListVersionsByFunction",
      "lambda:GetPolicy",
      "lambda:GetFunctionConfiguration"
    ]
    resources = ["*"]
  }

  statement {
    sid    = "Logs"
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
      "logs:DescribeLogStreams",
      "logs:DescribeLogGroups",
      "logs:ListTagsForResource"
    ]
    resources = ["*"]
  }

  statement {
    sid    = "PassRole"
    effect = "Allow"
    actions = [
      "iam:PassRole",
      "iam:GetRole",
      "iam:GetPolicy",
      "iam:ListRolePolicies",
      "iam:GetPolicyVersion",
      "iam:ListAttachedRolePolicies"
    ]
    resources = ["*"]
  }

  statement {
    sid = "S3TerraformState"
    effect = "Allow"
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:ListBucket"
    ]
    resources = [
      "arn:aws:s3:::yahlatci-terraform-states",
      "arn:aws:s3:::yahlatci-terraform-states/*"
    ]
  }

  statement {
    sid = "Ec2"
    effect = "Allow"
    actions = [
      "ec2:DescribeSubnets",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeSecurityGroupRules",

    ]
    resources = ["*"]
  }

  statement {
    sid = "ApplicationAutoScaling"
    effect = "Allow"
    actions = [
      "application-autoscaling:RegisterScalableTarget",
      "application-autoscaling:DeregisterScalableTarget",
      "application-autoscaling:DescribeScalableTargets",
      "application-autoscaling:PutScalingPolicy",
      "application-autoscaling:DeleteScalingPolicy",
      "application-autoscaling:DescribeScalingPolicies",
      "application-autoscaling:ListTagsForResource"
    ]
    resources = ["*"]
  }

  statement {
    sid = "ElasticLoadBalancing"
    effect = "Allow"
    actions = [
      "elasticloadbalancing:DescribeTargetHealth"
    ]
    resources = ["*"]
  }
}


resource "aws_iam_policy" "github_actions" {
  name   = "GithubActionsPolicy"
  policy = data.aws_iam_policy_document.github_permissions.json
}

resource "aws_iam_role_policy_attachment" "github_actions_attach" {
  role       = aws_iam_role.github_actions.name
  policy_arn = aws_iam_policy.github_actions.arn
}


output "github_oidc_provider_arn" {
  value = aws_iam_openid_connect_provider.github.arn
}

output "github_actions_role_arn" {
  value = aws_iam_role.github_actions.arn
}
