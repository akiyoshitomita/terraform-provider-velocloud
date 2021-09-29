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
}

data "velocloud_license_list" "license" {
  edition        = "PREMIUM"
  region         = "APAC"
  term_months    = 60
  bandwidth_tier = "001GW"
}

resource "velocloud_edge" "edge1" {
  name             = "tomita1"
  configuration_id = 5537
  model_number     = "virtual"
  edge_license_id  = data.velocloud_license_list.license.licenses[0].license_id
  custom_info      = "abc"
  site {
    name = "test2"
  }
}

#output "test1" {
#  value = velocloud_edge.edge1
#}

#output "test2"{
#  value = data.velocloud_license_list.license
#}

#output "a" {
#  value = data.velocloud_edge.edge0
#}
