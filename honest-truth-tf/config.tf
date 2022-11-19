provider "aws" {
  region = "us-east-1"
  profile = "honest_truth_admin"
}

terraform {
  required_version = ">= 1.2.0"

  backend "s3" {
    bucket = "honest-truth-terraform"
    key = "terraform.state"
    region = "us-east-1"
    profile = "honest_truth_admin"
  }

  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.16"
    }
  }
}
