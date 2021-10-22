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
    #name               = "Global Segment"
    #segment_logical_id = "2d36fc49-701d-4da8-9eb1-c32b5b82f7b7"
    segment_id = 0
    firewall_rule {
      source_type      = "vlan"
      source_vlan      = 1
      destination_type = "any"
      name             = "test2"
      source_port      = "80-82"
      logging          = true
    }
    firewall_rule {
      source_type      = "vlan"
      source_vlan      = 1
      destination_type = "any"
      name             = "test3"
      source_port      = "80-82"
      logging          = true
    }
  }
  port_forwarding_rule {
    name       = "rule1"
    protocol   = "tcp"
    interface  = "GE2"
    outside_ip = "192.168.1.100"
    wan_ports  = "80-443"
    lan_ip     = "192.168.2.100"
    lan_port   = 80
    segment_id = 0
  }

  nat_rule {
    name             = "nat1"
    outbound_ip      = "1.1.1.1"
    interface        = "GE2"
    inside_ip        = "192.168.1.100"
    segment_id       = 0
    outbound_traffic = true
    allow_protocol   = "all"

  }

  stateful_firewall {
    edge_overwrite              = true
    establieshd_tcp_timeout     = 7200
    non_established_tcp_timeout = 300
    udp_timeout                 = 300
    other_timeout               = 60
  }
  network_flood_protection{
    #edge_overwrite = true
    #new_connection_threshold = 25
  }

  edge_access{
    edge_overwrite = true
    ssh = true
    ssh_allow = ["1.1.1.1"]
    webui = true
    webui_allow = ["1.1.1.1"]
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
