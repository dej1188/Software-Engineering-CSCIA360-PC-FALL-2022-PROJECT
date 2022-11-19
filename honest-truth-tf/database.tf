resource "aws_db_parameter_group" "honest_truth" {
  name = "honest-truth"
  family = "postgres14"

  parameter {
    name = "log_connections"
    value = 1
  }
}

resource "aws_db_subnet_group" "honest_truth" {
  name = "honest-truth-db-subnet-groups"
  subnet_ids = [
    aws_subnet.honest_truth_subnet_private_a.id,
    aws_subnet.honest_truth_subnet_private_b.id
  ]
}

resource "aws_security_group" "honest_truth_db" {
  name   = "honest-truth-db-sg"
  vpc_id = aws_vpc.honest_truth_vpc.id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.honest_truth_vpc.cidr_block]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

data "aws_secretsmanager_secret_version" "honest_truth_db_password_db" {
  secret_id = aws_secretsmanager_secret.ht_database_password.id

  depends_on = [aws_secretsmanager_secret.ht_database_password]
}

resource "aws_db_instance" "honest_truth" {
  identifier = "honest-truth-db"
  allocated_storage = 10
  engine = var.db_engine
  engine_version = var.db_engine_version
  instance_class = var.db_instance_class
  db_name = var.db_name
  username = var.db_username
  password = data.aws_secretsmanager_secret_version.honest_truth_db_password_db.secret_string
  parameter_group_name = aws_db_parameter_group.honest_truth.name
  skip_final_snapshot = true
  backup_retention_period = 0
  apply_immediately = true
  db_subnet_group_name = aws_db_subnet_group.honest_truth.name
  vpc_security_group_ids = [
    aws_security_group.honest_truth_db.id
  ]
}