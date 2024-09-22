output "public_dns" {
  value = aws_instance.clock.public_dns
}

output "public_domain" {
  value = var.domain
}

output "public_ip" {
  value = aws_instance.clock.public_ip
}
