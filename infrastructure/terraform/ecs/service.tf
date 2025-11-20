module "ecs_service" {
  source  = "terraform-aws-modules/ecs/aws//modules/service"
  version = "v6.9.0"

  name                          = "sre-case-study-service"
  cpu                           = 1024
  memory                        = 2048
  autoscaling_max_capacity      = 4
  autoscaling_min_capacity      = 3
  availability_zone_rebalancing = "ENABLED"
  cluster_arn                   = module.ecs_cluster.arn
  create                        = true
  create_iam_role               = false
  iam_role_arn = aws_iam_role.ecs_task.arn
  create_tasks_iam_role = false
  tasks_iam_role_arn = aws_iam_role.ecs_task.arn
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
      referenced_security_group_id = module.alb.security_group_id
    }
  }
  deployment_minimum_healthy_percent = 100
  subnet_ids                         = module.vpc.private_subnets
  load_balancer = {
    service = {
      target_group_arn = aws_alb_target_group.api.arn
      container_name   = "sre-case-study-api"
      container_port   = 3000
    }
  }
  container_definitions = {
    api = {
      name      = "sre-case-study-api"
      image     = "887495603804.dkr.ecr.eu-west-1.amazonaws.com/sre-case-study:0c479c096d233e620be747ad1e8f8dbf7189f375"
      essential = true
      healthCheck = {
        command     = ["CMD-SHELL", "curl -f http://localhost:3000/health || exit 1"]
        interval    = 30
        timeout     = 5
        retries     = 3
        startPeriod = 10
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
          value = "local"
        }
      ]
    }
  }
}