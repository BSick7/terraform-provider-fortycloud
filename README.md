# terraform-provider-fortycloud

Terraform provider for Forty Cloud

## Usage

### Provider Configuration

Provider information can be set up in a *.tf file or through environment variables.

```
provider "fortycloud" {
	access_key = "access_key"
	secret_key = "secret_key"
}
```

```
FORTYCLOUD_ACCESS_KEY = "access_key"
FORTYCLOUD_SECRET_KEY = "secret_key"
```

## Install (download)

Download the binary for your platform from [releases](https://github.com/bsick7/terraform-provider-fortycloud/releases).

## Install (build)

You will need go1.6 and terraform 0.6.16 to build.

### Build

```bash
$ mkdir -p $GOPATH/src/github.com/bsick7
$ cd $GOPATH/src/github.com/bsick7
$ git clone git@github.com:bsick7/terraform-provider-fortycloud.git
$ make deps
$ make install
```

### How to use

See [examples](/examples) directory for use cases.
