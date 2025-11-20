module "ecs_cluster" {
  source  = "terraform-aws-modules/ecs/aws//modules/cluster"
  version = "v6.9.0"

  name = "sre-case-study"

  default_capacity_provider_strategy = {
    FARGATE_SPOT = {
      weight = 1
    }
  }

}

output "ecs_cluster_arn" {
  value = module.ecs_cluster.arn
  
}