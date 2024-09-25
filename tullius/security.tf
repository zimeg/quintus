# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/security_group
resource "aws_security_group" "hourglass" {
  name = "qt.hourglass"
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/vpc_security_group_egress_rule
resource "aws_vpc_security_group_egress_rule" "outbound" {
  security_group_id = aws_security_group.hourglass.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "-1"
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/vpc_security_group_ingress_rule
resource "aws_vpc_security_group_ingress_rule" "https" {
  security_group_id = aws_security_group.hourglass.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "tcp"
  from_port   = 443
  to_port     = 443
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/vpc_security_group_ingress_rule
resource "aws_vpc_security_group_ingress_rule" "ntp" {
  security_group_id = aws_security_group.hourglass.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "udp"
  from_port   = 123
  to_port     = 123
}
