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


#data "velocloud_edge" "edge0" {
#  name = "tomita_home"
#}

data "velocloud_license_list" "license" {
  edition        = "PREMIUM"
  region         = "APAC"
  term_months    = 60
  bandwidth_tier = "001GW"
}

#data "velocloud_profile" "profile" {
#  name = "tomita_test2"
#}

resource "velocloud_profile" "profile1" {
  name        = "tomtia_profile1"
  description = "this is description 2"
}

resource "velocloud_edge" "edge1" {
  name            = "tomita1"
  profile_id      = velocloud_profile.profile1.id
  model_number    = "virtual"
  edge_license_id = data.velocloud_license_list.license.licenses[0].license_id
  custom_info     = "abc"
  site {
    name = "test2"
  }
}


#output "test1" {
#  value = velocloud_edge.edge1
#}

#output "test2"{
#  value = data.velocloud_profile.profile
#}

#output "a" {
#  value = data.velocloud_edge.edge0
#}
