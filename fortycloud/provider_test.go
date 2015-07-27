package fortycloud

import (
	"log"
	"os"
	"testing"
	
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"fortycloud": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FORTYCLOUD_USERNAME"); v == "" {
		t.Fatal("FORTYCLOUD_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("FORTYCLOUD_PASSWORD"); v == "" {
		t.Fatal("FORTYCLOUD_PASSWORD must be set for acceptance tests")
	}
	if v := os.Getenv("FORTYCLOUD_TENANTNAME"); v == "" {
		t.Fatal("FORTYCLOUD_TENANTNAME must be set for acceptance tests")
	}
	if v := os.Getenv("FORTYCLOUD_FORMUSERNAME"); v == "" {
		t.Fatal("FORTYCLOUD_FORMUSERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("FORTYCLOUD_FORMPASSWORD"); v == "" {
		t.Fatal("FORTYCLOUD_FORMPASSWORD must be set for acceptance tests")
	}
}