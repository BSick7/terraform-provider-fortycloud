resource "fortycloud_gateway" "gw1" {
  public_ip = "54.165.11.200"
}

resource "fortycloud_gateway" "gw2" {
  public_ip = "54.172.64.69"
}

resource "fortycloud_ip_sec_connection" "gw1_gw2" {
  gateway_a = "${fortycloud_gateway.gw1.name}"
  gateway_b = "${fortycloud_gateway.gw2.name}"
}

resource "fortycloud_subnet" "subnet1" {
  name = "subnet1"
  description = "Subnet 1"
  cidr = "10.1.0.0/16"
  gateway_id = "${fortycloud_gateway.gw1.id}"
}

resource "fortycloud_subnet" "subnet2" {
  name = "subnet2"
  description = "Subnet 2"
  cidr = "10.2.0.0/16"
  gateway_id = "${fortycloud_gateway.gw2.id}"
}
