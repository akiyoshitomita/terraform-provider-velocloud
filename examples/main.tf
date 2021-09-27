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

#data "velocloud_edge" "edge0" {
#  name = "tomita_home"
#}

resource "velocloud_edge" "edge1" {
  name             = "tomita1"
  configuration_id = 5537
  model_number     = "virtual"
  edge_license_id  = 175
  custom_info      = "abc"
  site {
    name = "test2"
  }
}

output "test1" {
  value = velocloud_edge.edge1
}

#output "a" {
#  value = data.velocloud_edge.edge0
#}
