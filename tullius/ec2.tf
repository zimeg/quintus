# https://search.opentofu.org/provider/opentofu/random/latest/docs/resources/pet
resource "random_pet" "cicero" {
  length = 2
}

# https://search.opentofu.org/provider/hashicorp/aws/latest/docs/resources/ami
resource "aws_ami" "cicero" {
  name                = "ciceroami-${random_pet.cicero.id}"
  virtualization_type = "hvm"
  ena_support         = true
  root_device_name    = "/dev/xvda"

  ebs_block_device {
    device_name = "/dev/xvda"
    snapshot_id = aws_ebs_snapshot_import.cicero.id
  }
}

# https://search.opentofu.org/provider/hashicorp/aws/latest/docs/resources/instance
resource "aws_instance" "clock" {
  ami                    = aws_ami.cicero.id
  instance_type          = "t3.nano"
  vpc_security_group_ids = [aws_security_group.hourglass.id]

  lifecycle {
    create_before_destroy = true
  }
}
