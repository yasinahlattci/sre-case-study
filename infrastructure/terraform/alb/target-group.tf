resource "aws_alb_target_group" "api" {
  name        = "api-tg"
  port        = 3000
  protocol    = "HTTP"
  vpc_id      = data.terraform_remote_state.vpc.outputs.vpc_id
  target_type = "ip"
  health_check {
    path                = "/health"
    protocol            = "HTTP"
    matcher             = "200"
    port                = "3000"
    interval            = 15
    timeout             = 3
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }
  
  deregistration_delay = 30
}

resource "aws_alb_target_group" "lambda" {
  name        = "lambda-tg"
  vpc_id      = data.terraform_remote_state.vpc.outputs.vpc_id
  target_type = "lambda"
}
