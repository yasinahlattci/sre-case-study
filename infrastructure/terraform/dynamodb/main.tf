resource "aws_dynamodb_table" "default" {
    name = "picusv4"
    hash_key = "objectID"

    attribute {
      name = "objectID"
      type = "S"
    }
    billing_mode = "PAY_PER_REQUEST"
}


output "dynamodb_table_name" {
    value = aws_dynamodb_table.default.name
  
}