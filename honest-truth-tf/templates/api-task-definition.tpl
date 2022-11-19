[
    {
      "name": "${task_name}",
      "image": "${image_url}",
      "portMappings": [
        {
          "containerPort": ${port}
        }
      ],
      "environment": [
        {
            "name": "DB_HOST",
            "value": "${db_host}"
        },
        {
            "name": "DB_PORT",
            "value": "${db_port}"
        },
        {
            "name": "DB_USER",
            "value": "${db_user}"
        },
        {
            "name": "DB_NAME",
            "value": "${db_name}"
        },
        {
            "name": "DB_AUTO_MIGRATE",
            "value": "${auto_migrate}"
        },
        {
            "name": "ENV",
            "value": "${environment}"
        },
        {
            "name": "PORT",
            "value": "${port}"
        }
      ],
      "secrets": [
        {
            "name": "DB_PASS",
            "valueFrom": "${db_pass}"
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-region": "${region}",
          "awslogs-group": "${log_group}",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
]