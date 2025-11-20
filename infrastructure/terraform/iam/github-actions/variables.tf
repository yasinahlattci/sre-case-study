variable "region" {
  type    = string
  default = "eu-west-1"
}

variable "github_username" {
  type        = string
  description = "GitHub username"
  default = "yasinahlattci"
}

variable "github_repo" {
  type        = string
  description = "GitHub repo name"
  default = "sre-case-study"
}