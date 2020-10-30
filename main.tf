terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "us-east-1"
  # Credentials are not configured. Env variables should be correctly configured.
}

variable "s3_bucket" {
    type = string
    default = "mporracindie-tf-test"
}

resource "aws_s3_bucket" "b" {
  bucket = var.s3_bucket
  acl    = "private"
}

variable "s3_files" {
    type = list
    default = ["test1.txt","test2.txt"]
}

resource "aws_s3_bucket_object" "object" {
  bucket = var.s3_bucket
  depends_on = [aws_s3_bucket.b]
  count = length(var.s3_files)
  key    = var.s3_files[count.index]
  content = timestamp()
}

output "bucket_info" {
  value = aws_s3_bucket.b.id
}

output "files_info" {
  value = aws_s3_bucket_object.object.*.id
}

