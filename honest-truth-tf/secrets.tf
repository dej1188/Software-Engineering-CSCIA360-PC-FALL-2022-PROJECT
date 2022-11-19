resource "random_password" "ht_database_password" {
  length = 16
  special = true
  override_special = "_!%^"
}

resource "aws_secretsmanager_secret" "ht_database_password" {
  name = "ht-db-pass"
  recovery_window_in_days = 0
}

resource "aws_secretsmanager_secret_version" "ht_database_password" {
  secret_id = aws_secretsmanager_secret.ht_database_password.id
  secret_string = random_password.ht_database_password.result
}

resource "aws_iam_role_policy" "password_policy_secretsmanager" {
  name = "password-policy-secretsmanager"
  role = aws_iam_role.honest_truth_api_task_execution_role.id

  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "secretsmanager:GetSecretValue"
        ],
        "Effect": "Allow",
        "Resource": [
          "${aws_secretsmanager_secret.ht_database_password.arn}"
        ]
      }
    ]
  }
  EOF
}