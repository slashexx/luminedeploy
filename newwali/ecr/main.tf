
resource "aws_ecr_repository" "lavi-ki-repo" {
  name = "lavi-ki-repo"
  image_tag_mutability = "MUTABLE"
  image_scanning_configuration {
    scan_on_push = true
  }
}
