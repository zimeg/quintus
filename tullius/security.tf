# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/security_group
resource "aws_security_group" "hourglass" {
  name = "qt.hourglass"
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/vpc_security_group_egress_rule
resource "aws_vpc_security_group_egress_rule" "outbound" {
  security_group_id = aws_security_group.hourglass.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "-1"
  from_port   = 0
  to_port     = 0
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/vpc_security_group_ingress_rule
resource "aws_vpc_security_group_ingress_rule" "ssh" {
  security_group_id = aws_security_group.hourglass.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "tcp"
  from_port   = 22
  to_port     = 22
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/vpc_security_group_ingress_rule
resource "aws_vpc_security_group_ingress_rule" "ntp" {
  security_group_id = aws_security_group.hourglass.id

  cidr_ipv4   = "0.0.0.0/0"
  ip_protocol = "udp"
  from_port   = 123
  to_port     = 123
}

# https://search.opentofu.org/provider/opentofu/tls/latest/docs/resources/private_key
resource "tls_private_key" "state_ssh_key" {
  algorithm = "RSA"
}

# https://search.opentofu.org/provider/hashicorp/local/latest/docs/resources/sensitive_file
resource "local_sensitive_file" "machine_ssh_key" {
  content         = tls_private_key.state_ssh_key.private_key_pem
  filename        = "${path.module}/id_rsa.pem"
  file_permission = "0600"
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/datasources/key_pair
resource "aws_key_pair" "generated_key" {
  key_name   = "generated-key-${sha256(tls_private_key.state_ssh_key.public_key_openssh)}"
  public_key = tls_private_key.state_ssh_key.public_key_openssh
}
