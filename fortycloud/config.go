package fortycloud

import (
    "log"
    
    "github.com/mdl/fortycloud-sdk-go/api"
)

type Config struct {
    Username      string
    Password      string
    TenantName    string
    FormUsername  string
    FormPassword  string
}

// Client() returns a new Api for accessing FortyCloud.
func (config *Config) Api() (*fortycloud.Api, error) {
    api := fortycloud.NewApi()
    api.SetApiCredentials(config.Username, config.Password, config.TenantName)
    api.SetFormsCredentials(config.FormUsername, config.FormPassword)
    log.Printf("[INFO] Forty Cloud Client configured for users: %s, %s", config.Username, config.FormUsername)
    return api, nil
}