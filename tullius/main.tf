terraform {
  required_version = "1.8.2"

  required_providers {
    # https://search.opentofu.org/provider/opentofu/aws/latest
    aws = {
      source  = "opentofu/aws"
      version = "5.66.0"
    }
  }

  backend "s3" {
    bucket         = "architectf"
    key            = "quintus"
    region         = "us-east-1"
    dynamodb_table = "architectf-timeline"
  }
}

provider "aws" {
  region = "us-east-1"
}
