package fortycloud

import (
    "log"
    "net/http"
    
    "github.com/mdl/fortycloud-sdk-go/api"
)

type Config struct {
    User          string
    Password      string
}

// Client() returns a new Service for accessing FortyCloud.
func (c *Config) Service() (*fortycloud.Service, error) {
    /*
    service := fortycloud.NewService(&http.Client{
        Transport: &fortycloud.Transport{
            Username: c.User,
            Password: c.Password
        }
    })
    */
    
    service := fortycloud.NewService(&http.Client{
        
    })
    
    log.Printf("[INFO] Forty Cloud Client configured for user: %s", c.User)
    
    return service, nil
}