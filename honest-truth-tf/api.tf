resource "aws_ecr_repository" "honest_truth_api" {
  name = "honest-truth-api"
  force_delete = true

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecs_cluster" "honest_truth_api" {
  name = "honest-truth-api"
}

resource "aws_cloudwatch_log_group" "honest_truth_api" {
  name = "/ecs/honest-truth-api"
}

resource "aws_ecs_service" "honest_truth_api" {
  name = "honest-truth-api"
  task_definition = aws_ecs_task_definition.honest_truth_api.arn
  cluster = aws_ecs_cluster.honest_truth_api.id
  launch_type = "FARGATE"

  desired_count = 1

  load_balancer {
    target_group_arn = aws_lb_target_group.honest_truth_api.arn
    container_name = "honest-truth-api"
    container_port = 80
  }

  network_configuration {
    assign_public_ip = false

    security_groups = [
      aws_security_group.honest_truth_http_80.id
    ]

    subnets = [
      aws_subnet.honest_truth_subnet_private_a.id,
      aws_subnet.honest_truth_subnet_private_b.id
    ]
  }
}


data "template_file" "honest_truth_api_task_definition" {
  template = file("${path.module}/templates/api-task-definition.tpl")

  vars = {
    task_name = "honest-truth-api"
    image_url = "${aws_ecr_repository.honest_truth_api.repository_url}:latest"
    region = "us-east-1"
    log_group = aws_cloudwatch_log_group.honest_truth_api.name
    db_host = aws_db_instance.honest_truth.address
    db_port = aws_db_instance.honest_truth.port
    db_user = var.db_username
    db_pass = aws_secretsmanager_secret.ht_database_password.arn
    db_name = var.db_name
    environment = var.environment
    port = "80"
    auto_migrate = "1"
  }
}

resource "aws_ecs_task_definition" "honest_truth_api" {
  family = "honest-truth-api"

  container_definitions = data.template_file.honest_truth_api_task_definition.rendered
  execution_role_arn = aws_iam_role.honest_truth_api_task_execution_role.arn

  cpu = var.api_cpu
  memory = var.api_memory
  requires_compatibilities = ["FARGATE"]

  network_mode = "awsvpc"
}

resource "aws_iam_role" "honest_truth_api_task_execution_role" {
  name               = "honest-truth-task-execution-role"
  assume_role_policy = data.aws_iam_policy_document.ecs_task_assume_role.json
}

data "aws_iam_policy_document" "ecs_task_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

# Normally we'd prefer not to hardcode an ARN in our Terraform, but since this is an AWS-managed
# policy, it's okay.
data "aws_iam_policy" "ecs_task_execution_role" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# Attach the above policy to the execution role.
resource "aws_iam_role_policy_attachment" "ecs_task_execution_role" {
  role       = aws_iam_role.honest_truth_api_task_execution_role.name
  policy_arn = data.aws_iam_policy.ecs_task_execution_role.arn
}

resource "aws_lb_target_group" "honest_truth_api" {
  name        = "honest-truth-api"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.honest_truth_vpc.id

  health_check {
    path    = "/health"
    interval = 60
  }

  depends_on = [aws_alb.honest_truth_api]
}

resource "aws_security_group" "honest_truth_http_80" {
  name        = "http"
  description = "HTTP traffic"
  vpc_id      = aws_vpc.honest_truth_vpc.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "TCP"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_alb" "honest_truth_api" {
  name               = "honest-truth-api-lb"
  internal           = false
  load_balancer_type = "application"

  subnets = [
    aws_subnet.honest_truth_subnet_public_a.id,
    aws_subnet.honest_truth_subnet_public_b.id
  ]

  security_groups = [
    aws_security_group.honest_truth_http_80.id
  ]

  depends_on = [aws_internet_gateway.honest_truth_igw]
}

resource "aws_alb_listener" "honest_truth_api_http" {
  load_balancer_arn = aws_alb.honest_truth_api.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.honest_truth_api.arn
  }
}