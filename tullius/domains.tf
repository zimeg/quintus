# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/route53_zone
resource "aws_route53_zone" "timekeeper" {
  name = var.domain
}
