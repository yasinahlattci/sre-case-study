module "alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "v10.2.0"

  name = "sre-case-study-alb"
  access_logs = {
    bucket  = ""
    enabled = false
  }
  create                = true
  create_security_group = true
  security_group_name  = "sre-case-study-alb-sg"
  security_group_ingress_rules = {
    allow_yasin = {
      name      = "allow-yasin"
      cidr_ipv4 = "178.233.42.39/32"
      from_port = "80"
      to_port   = "80"
    }
  }

  security_group_egress_rules = {
    all_vpc = { 
      ip_protocol    = "-1"
      cidr_ipv4     = module.vpc.vpc_cidr_block
    }

  }
  default_port          = "80"
  default_protocol      = "HTTP"
  internal              = false
  ip_address_type       = "ipv4"
  load_balancer_type    = "application"
  region                = var.region
  listeners = {
    http = {
      port     = "80"
      protocol = "HTTP"
      fixed_response = {
        content_type = "text/plain"
        message_body = "404: page not found"
        status_code  = "404"
      }
      rules = {
        picus-get = {
          listener_key     = "http"
          priority         = 10
          target_group_key = "api"
          conditions = [{
            path_pattern = {
              values = ["/picus/get/*", "/picus/list", "/health"]
            }
            },
            {
              http_request_method = {
                values = ["GET"]
              }
            }
          ]
          actions = [
                      {
              forward = {
                target_group_arn = aws_alb_target_group.api.arn
              }
            }

          ]
        }
        picus-put = {
          listener_key     = "http"
          priority         = 20
          target_group_key = "api"
          conditions = [{
            path_pattern = {
              values = ["/picus/put"]
            }
            },
            {
              http_request_method = {
                values = ["POST"]
              }
            }
          ]
          actions = [
            {
              forward = {
                target_group_arn = aws_alb_target_group.api.arn
              }
            }
          ]
        }
        picus-delete = {
          listener_key     = "http"
          priority         = 30
          target_group_key = "api"
          conditions = [{
            path_pattern = {
              values = ["/picus/*"]
            }
            },
            {
              http_request_method = {
                values = ["DELETE"]
              }
            }
          ]
          actions = [
            {
              fixed_response = {
                content_type = "text/plain"
                message_body = "404: DELETE page not found"
                status_code  = "404"
              }
            }
          ]
        }
      }
    }
  }
  vpc_id              = module.vpc.vpc_id
  subnets             = module.vpc.public_subnets
}