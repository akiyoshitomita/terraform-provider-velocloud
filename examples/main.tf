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

resource "velocloud_firewall" "edge_firewall" {
  edge_id                                   = velocloud_edge.edge1.edge_id
  firewall_enabled                          = true
  edge_overwrite_statefull_firewall_enabled = true
  edge_overwrite_syslog_forwarding          = true
  statefull_firewall_enabled                = true

  segments {
    name               = "Global Segment"
    #segment_logical_id = "2d36fc49-701d-4da8-9eb1-c32b5b82f7b7"
    firewall_rule {
      #action = "allow"
      source_type = "vlan"
      source_vlan = 1
      destination_type = "any"
      name        = "test2"
      source_mac  = "aa:bb:cc:dd:ee:ff"
      source_port = "80-82"
      logging = true 
    }
  }
#  segments {
#    name               = "test"
#    segment_logical_id = "683e5aed-dc0b-47e4-bf66-ec7842e37dee"
#  }
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
