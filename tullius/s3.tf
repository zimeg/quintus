# https://search.opentofu.org/provider/opentofu/aws/latest/docs/datasources/s3_bucket
resource "aws_s3_bucket" "server" {
  bucket = var.vhs
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/s3_bucket_ownership_controls
resource "aws_s3_bucket_ownership_controls" "vhs" {
  bucket = aws_s3_bucket.server.id

  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/s3_bucket_public_access_block
resource "aws_s3_bucket_public_access_block" "lock" {
  bucket = aws_s3_bucket.server.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# https://search.opentofu.org/provider/opentofu/aws/latest/docs/datasources/s3_object
resource "aws_s3_object" "upload" {
  bucket = aws_s3_bucket.server.id
  key    = "cicero.vhd"
  source = var.image
}
