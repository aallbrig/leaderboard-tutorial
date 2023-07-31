# run a cloudformation stack
resource "aws_cloudformation_stack" "stack" {
  name = "leaderboard-stack"
  template_body = file("../cf/stack.yaml")
}