terraform {
  required_version = "1.10.6"

  required_providers {
    # https://search.opentofu.org/provider/hashicorp/aws/latest
    aws = {
      source  = "hashicorp/aws"
      version = "6.15.0"
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
