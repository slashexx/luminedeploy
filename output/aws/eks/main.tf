
resource "aws_eks_cluster" "bhen-kichut" {
  name     = "bhen-kichut"
  role_arn = "arn:aws:iam::your_account_id:role/eks-service-role"

  vpc_config {
    subnet_ids = ["subnet-0123456789abcdef0", "subnet-abcdef0123456789"]
  }
}
