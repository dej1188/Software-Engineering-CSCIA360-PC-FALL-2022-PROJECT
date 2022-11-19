resource "aws_vpc" "honest_truth_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "honest_truth_subnet_public_a" {
  vpc_id = aws_vpc.honest_truth_vpc.id
  cidr_block = "10.0.1.0/25"
  availability_zone = "us-east-1c"
  tags = {
    "Name" = "public | us-east-1c"
  }
}

resource "aws_subnet" "honest_truth_subnet_public_b" {
  vpc_id = aws_vpc.honest_truth_vpc.id
  cidr_block = "10.0.1.128/25"
  availability_zone = "us-east-1d"
  tags = {
    "Name" = "public | us-east-1d"
  }
}

resource "aws_subnet" "honest_truth_subnet_private_a" {
  vpc_id = aws_vpc.honest_truth_vpc.id
  cidr_block = "10.0.2.0/25"
  availability_zone = "us-east-1c"
  tags = {
    "Name" = "private | us-east-1c"
  }
}

resource "aws_subnet" "honest_truth_subnet_private_b" {
  vpc_id = aws_vpc.honest_truth_vpc.id
  cidr_block = "10.0.2.128/25"
  availability_zone = "us-east-1d"
  tags = {
    "Name" = "private | us-east-1d"
  }
}

resource "aws_route_table" "honest_truth_public_rt" {
  vpc_id = aws_vpc.honest_truth_vpc.id
}

resource "aws_route_table_association" "honest_truth_public_a" {
  subnet_id = aws_subnet.honest_truth_subnet_public_a.id
  route_table_id = aws_route_table.honest_truth_public_rt.id
}

resource "aws_route_table_association" "honest_truth_public_b" {
  subnet_id = aws_subnet.honest_truth_subnet_public_b.id
  route_table_id = aws_route_table.honest_truth_public_rt.id
}

resource "aws_route_table" "honest_truth_private_rt" {
  vpc_id = aws_vpc.honest_truth_vpc.id
}

resource "aws_route_table_association" "honest_truth_private_a" {
  subnet_id = aws_subnet.honest_truth_subnet_private_a.id
  route_table_id = aws_route_table.honest_truth_private_rt.id
}

resource "aws_route_table_association" "honest_truth_private_b" {
  subnet_id = aws_subnet.honest_truth_subnet_private_b.id
  route_table_id = aws_route_table.honest_truth_private_rt.id
}

resource "aws_eip" "honest_truth_nat" {
  vpc = true
}

resource "aws_nat_gateway" "honest_truth_ngw" {
  subnet_id = aws_subnet.honest_truth_subnet_public_a.id
  allocation_id = aws_eip.honest_truth_nat.id

  depends_on = [aws_internet_gateway.honest_truth_igw]
}

resource "aws_route" "honest_truth_private_ngw_route" {
  route_table_id = aws_route_table.honest_truth_private_rt.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id = aws_nat_gateway.honest_truth_ngw.id
}

resource "aws_internet_gateway" "honest_truth_igw" {
  vpc_id = aws_vpc.honest_truth_vpc.id
}

resource "aws_route" "honest_truth_public_igw_route" {
  route_table_id = aws_route_table.honest_truth_public_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id = aws_internet_gateway.honest_truth_igw.id
}
