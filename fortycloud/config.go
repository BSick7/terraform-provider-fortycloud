package fortycloud

import (
	fc "github.com/bsick7/fortycloud-sdk-go/api"
	"log"
)

type Config struct {
	AccessKey string
	SecretKey string
}

// Client() returns a new Api for accessing FortyCloud.
func (config *Config) Api() (*fc.Api, error) {
	api := fc.NewApi(nil)
	api.SetAccessCredentials(config.AccessKey, config.SecretKey)
	log.Printf("[INFO] Forty Cloud Client configured.")
	return api, nil
}
