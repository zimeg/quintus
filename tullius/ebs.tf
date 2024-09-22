# https://search.opentofu.org/provider/opentofu/aws/latest/docs/resources/ebs_snapshot_import
resource "aws_ebs_snapshot_import" "cicero" {
  role_name = aws_iam_role.vm.id

  disk_container {
    format = "VHD"

    user_bucket {
      s3_bucket = aws_s3_bucket.server.id
      s3_key    = aws_s3_object.upload.id
    }
  }

  lifecycle {
    replace_triggered_by = [
      aws_s3_object.upload
    ]
  }
}
