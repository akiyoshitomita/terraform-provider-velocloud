variable "vco" {
  type        = string
  description = "The address of vco. (ex. https://vco.local"
}

variable "username" {
  type        = string
  description = "user name of vco"
}

variable "password" {
  type      = string
  sensitive = true
}
