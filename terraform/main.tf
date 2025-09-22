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
  public_key = file("personal-site-key.pub")
}

resource "aws_instance" "go_app" {
  ami             = "ami-069e78ac413158ee1" # Amazon Linux 2023
  instance_type   = "t3.micro"
  key_name        = aws_key_pair.go_key.key_name
  security_groups = [aws_security_group.go_app_sg.name]

  tags = {
    Name = "personal-site"
  }

  lifecycle {
    create_before_destroy = true
  }

  user_data = <<-EOF
    #!/bin/bash
    # Update system
    sudo dnf update -y

    sudo dnf install -y git

    sudo dnf install -y docker
    sudo systemctl enable docker
    sudo systemctl start docker

    sudo usermod -aG docker ec2-user

    sudo curl -L "https://github.com/docker/compose/releases/download/v2.28.2/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose
  EOF

}


resource "aws_eip" "go_app_ip" {
  instance = aws_instance.go_app.id
}
