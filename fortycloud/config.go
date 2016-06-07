package fortycloud

import (
	fc "github.com/BSick7/fortycloud-sdk-go/api"
	"log"
)

type Config struct {
	AccessKey          string
	SecretKey          string
	FindGatewayTimeout string
}

// Client() returns a new Api for accessing FortyCloud.
func (config *Config) Api() (*fc.Api, error) {
	api := fc.NewApi(nil)
	api.SetAccessCredentials(config.AccessKey, config.SecretKey)
	if config.FindGatewayTimeout != "" {
		api.SetFindGatewayTimeout(config.FindGatewayTimeout)
	}
	log.Printf("[INFO] Forty Cloud Client configured.")
	return api, nil
}
