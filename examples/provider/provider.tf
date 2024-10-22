terraform {
  required_providers {
    lambda = {
      source  = "squat/lambda"
      version = "0.1.3"
    }
  }
}

provider "lambda" {
  # Configuration options
}