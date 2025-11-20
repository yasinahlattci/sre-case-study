resource "aws_ecr_repository" "lambda" {
  name = "sre-case-study-lambda"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = false
  }
}

resource "aws_ecr_repository" "api" {
  name = "sre-case-study-api"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = false
  }
}