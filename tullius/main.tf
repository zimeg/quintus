terraform {
  required_version = "1.8.2"

  required_providers {
    # https://search.opentofu.org/provider/opentofu/archive/latest
    archive = {
      source  = "opentofu/archive"
      version = "2.6.0"
    }
    # https://search.opentofu.org/provider/opentofu/aws/latest
    aws = {
      source  = "opentofu/aws"
      version = "5.66.0"
    }
    # https://search.opentofu.org/provider/opentofu/local/latest
    local = {
      source  = "opentofu/local"
      version = "2.5.1"
    }
    # https://search.opentofu.org/provider/opentofu/null/latest
    null = {
      source  = "opentofu/null"
      version = "3.2.2"
    }
    # https://search.opentofu.org/provider/opentofu/random/latest
    random = {
      source  = "opentofu/random"
      version = "3.6.2"
    }
    # https://search.opentofu.org/provider/opentofu/tls/latest
    tls = {
      source  = "opentofu/tls"
      version = "4.0.5"
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
