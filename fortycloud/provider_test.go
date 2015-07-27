package fortycloud

import (
	"log"
	"os"
	"testing"
	
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testProviders map[string]terraform.ResourceProvider
var testProvider *schema.Provider

func init() {
	testProvider = Provider().(*schema.Provider)
	testProviders = map[string]terraform.ResourceProvider{
		"fortycloud": testProvider,
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

func testPreCheck(t *testing.T) {
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