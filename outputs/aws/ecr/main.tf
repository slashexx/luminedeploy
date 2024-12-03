
resource "aws_ecr_repository" "hello-new-repo" {
  name = "hello-new-repo"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}
