resource "fortycloud_subnet" "subnet1" {
	name = "subnet1"
	description = "Subnet 1"
	subnet = "10.1.0.0/17"
}

resource "fortycloud_subnet" "subnet2" {
	name = "subnet2"
	description = "Subnet 2"
	subnet = "10.1.0.0/17"
}