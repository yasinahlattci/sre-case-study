module "ecs_service" {
  source  = "terraform-aws-modules/ecs/aws//modules/service"
  version = "v6.9.0"

  name                          = "${var.name}-service"
  cpu                           = 1024
  memory                        = 2048
  autoscaling_max_capacity      = 4
  autoscaling_min_capacity      = 3
  availability_zone_rebalancing = "ENABLED"
  cluster_arn                   = var.ecs_cluster_arn
  create                        = true
  create_iam_role               = false
  iam_role_arn = var.task_role_arn
  create_tasks_iam_role = false
  tasks_iam_role_arn = var.task_role_arn
  create_task_exec_iam_role = true
  task_exec_iam_role_name = "ScsEcsTaskExecutionRole"
  desired_count                 = 3
  region                        = var.region
  scheduling_strategy           = "REPLICA"
  wait_until_stable             = true
  security_group_egress_rules = {
    allow_all = {
      ip_protocol    = "-1"
      cidr_ipv4     = "0.0.0.0/0"
    }
  }
  security_group_ingress_rules = {
    allow_alb = {
      ip_protocol      = "-1"
      protocol        = "tcp"
      referenced_security_group_id = var.alb_security_group_id
    }
  }
  deployment_minimum_healthy_percent = 100
  subnet_ids                         = var.vpc_private_subnet_ids
  load_balancer = {
    service = {
      target_group_arn = var.target_group_arn
      container_name   = "sre-case-study-api"
      container_port   = 3000
    }
  }
  container_definitions = {
    api = {
      name      = "sre-case-study-api"
      image     = var.image_uri
      essential = true
      healthCheck = {
        command     = ["CMD-SHELL", "curl -f http://localhost:3000/health || exit 1"]
        interval    = 30
        timeout     = 5
        retries     = 3
        startPeriod = 5
      }
      portMappings = [
        {
          containerPort = 3000
          protocol      = "tcp"
        }
      ]
      environment = [
        {
          name  = "APP_ENV"
          value = var.app_env
        }
      ]
    }
  }
}