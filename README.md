# terraform-provider-fortycloud
Terraform provider for Forty Cloud

## Usage

### Provider Configuration
```
provider "fortycloud" {
	username = "username"
	password = "Passw0rd"
	tenantName = "tenant"
	formUsername = "formusername"
	formPassword = "formpassword"
}
```

### How to install

1. Download and install [Go](https://golang.org/doc/install). (follow instructions carefully)
2. Download and extract contents to location [terraform](https://terraform.io/downloads.html).
3. Place terraform directory on `PATH`.
4. Build plugin and place with terraform binaries.

```
$ cd $GOPATH
$ mdkir mdl && cd mdl
$ git clone git@github.com/BSick7/terraform-provider-fortycloud
$ go get
$ go build && cp ./terraform-provider-fortycloud <insert terraform directory>
```
