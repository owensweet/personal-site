provider "aws" {
  region = "ca-west-1"
}

resource "aws_security_group" "go_app_sg" {
  name        = "go-app-sg"
  description = "Allow SSH and HTTP"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_key_pair" "go_key" {
  key_name   = "personal-site-key"
  public_key = file("~/.ssh/personal-site-key.pub")
}

resource "aws_instance" "go_app" {
  ami             = "ami-069e78ac413158ee1" # Amazon Linux 2023
  instance_type   = "t3.micro"
  key_name        = aws_key_pair.go_key.key_name
  security_groups = [aws_security_group.go_app_sg.name]

  tags = {
    Name = "personal-site"
  }

  user_data = <<-EOF
    #!/bin/bash
    # Update system
    sudo dnf update -y

    sudo dnf install -y docker
    sudo systemctl enable docker
    sudo systemctl start docker

    sudo dnf install -y docker-compose-plugin

    sudo usermod -aG docker ec2-user

    sleep 10

    cd /home/ec2-user
    git clone https://github.com/owensweet/personal-site.git
    cd personal-site

    docker compose up -d
  EOF
}


resource "aws_eip" "go_app_ip" {
  instance = aws_instance.go_app.id
}
