# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/iam_role_policy_attachment
resource "aws_iam_role_policy_attachment" "vm" {
  role       = aws_iam_role.vm.id
  policy_arn = aws_iam_policy.vm.arn
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/iam_role
resource "aws_iam_role" "vm" {
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect    = "Allow"
        Principal = { Service = "vmie.amazonaws.com" }
        Action    = "sts:AssumeRole"
        Condition = {
          StringEquals = {
            "sts:Externalid" = "vmimport"
          }
        }
      }
    ]
  })
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/iam_role_policy
resource "aws_iam_policy" "vm" {
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:GetBucketLocation",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",
          "s3:GetBucketAcl"
        ]
        Resource = [
          "arn:aws:s3:::${aws_s3_bucket.server.id}",
          "arn:aws:s3:::${aws_s3_bucket.server.id}/*"
        ]
      },
      {
        Effect = "Allow"
        Action = [
          "ec2:ModifySnapshotAttribute",
          "ec2:CopySnapshot",
          "ec2:RegisterImage",
          "ec2:Describe*"
        ],
        Resource = "*"
      }
    ]
  })
}
