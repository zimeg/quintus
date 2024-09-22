# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/route53_zone
resource "aws_route53_zone" "timekeeper" {
  name = var.domain
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/route53_record
resource "aws_route53_record" "servius" {
  name    = var.domain
  type    = "A"
  zone_id = aws_route53_zone.timekeeper.id
  ttl     = 300

  records = [aws_instance.clock.public_ip]
}
