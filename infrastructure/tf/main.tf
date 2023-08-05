data "aws_vpc" "default" {
  default = true
}
data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

resource "aws_cloudformation_stack" "database_stack" {
  name = "leaderboard-database-stack"
  template_body = file("../cf/stack.yaml")
  parameters = {
    VpcId = data.aws_vpc.default.id
    SubnetIds = join(",", data.aws_subnets.default.ids)
  }
}

resource "aws_cloudformation_stack" "cache_stack" {
  name = "leaderboard-cache-stack"
  template_body = file("../cf/cache.yaml")
  parameters = {
    SubnetIds = data.aws_subnets.default.ids[0]
  }
}
