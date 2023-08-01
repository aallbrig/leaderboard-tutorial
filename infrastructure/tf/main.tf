data "aws_vpc" "default" {
  default = true
}
data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
}

resource "aws_cloudformation_stack" "stack" {
  name = "leaderboard-stack"
  template_body = file("../cf/stack.yaml")
  parameters = {
    VpcId = data.aws_vpc.default.id
    SubnetIds = join(",", data.aws_subnets.default.ids)
  }
}