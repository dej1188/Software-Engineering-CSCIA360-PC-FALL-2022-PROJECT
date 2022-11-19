resource "aws_route53_zone" "goodcodelabs_com" {
  name = "goodcodelabs.com"

}

resource "aws_route53_record" "goodcodelabs_com_honest_truth_api" {
  name = "honest-truth-api.goodcodelabs.com"
  zone_id = aws_route53_zone.goodcodelabs_com.id
  type = "A"

  alias {
    evaluate_target_health = false
    name = aws_alb.honest_truth_api.dns_name
    zone_id = aws_alb.honest_truth_api.zone_id
  }
}

#resource "aws_route53_record" "goodcodelabs_com_honest_truth_ui" {
#  name = "honest-truth.goodcodelabs.com"
#  zone_id = aws_route53_zone.goodcodelabs_com.id
#  type = "A"
#
#  alias {
#    evaluate_target_health = false
#    name = "${aws_s3_bucket.honest_truth_ui.bucket}.${aws_s3_bucket_website_configuration.honest_truth_ui.website_domain}"
#    zone_id = aws_s3_bucket.honest_truth_ui.hosted_zone_id
#  }
#}