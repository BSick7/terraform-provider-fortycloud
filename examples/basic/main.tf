resource "fortycloud_subnet" "subnet1" {
	name = "subnet1"
	description = "Subnet 1"
	subnet = "10.1.0.0/16"
}

resource "fortycloud_subnet" "subnet2" {
	name = "subnet2"
	description = "Subnet 2"
	subnet = "10.2.0.0/16"
}

resource "fortycloud_node" "node1" {
	public_ip = "54.165.11.200"
}

resource "fortycloud_node" "node2" {
	public_ip = "54.172.64.69"
}

resource "fortycloud_connection" "node1_node2" {
	peer_a_id = "${fortycloud_node.node1.peer_id}"
	peer_b_id = "${fortycloud_node.node2.peer_id}"
}