terraform {
  required_providers {
    velocloud = {
      version = "0.1"
      source  = "akiyoshitomita/edu/velocloud"
    }
  }
}

provider "velocloud" {
  vco      = var.vco
  username = var.username
  password = var.password
  //username = "aaabbb"
}

data "velocloud_order" "aaa" {
  id = 1
}

output "a" {
  value = data.velocloud_order.aaa
}
