resource "aws_s3_bucket" "honest_truth_ui" {
  bucket = "honest-truth-ui"
  force_destroy = "true"
}

resource "aws_s3_bucket_website_configuration" "honest_truth_ui" {
  bucket = aws_s3_bucket.honest_truth_ui.id

  index_document {
    suffix = "index.html"
  }
}

resource "aws_s3_bucket_acl" "honest_truth_ui" {
  bucket = aws_s3_bucket.honest_truth_ui.id
  acl = "public-read"
}