package config

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "fed"

type config struct {

	// Account info
	User string `required:"true" description:"Username for fed2 account"`
	Pass string `required:"true" description:"Password for fed2 account"`
	Host string `default:"play.federation2.com" description:"Hostname for the fed2 service"`
	Port string `default:"30003" required:"true" description:"Port for the fed2 service"`

	// R.Aft config
	BotMode string `default:"wander" description:"Mode for the bot (ex: wandering, mining, ...)"`

	// internals
	confid uuid.UUID
}

func (c *config) String() string {
	format := "Account {\n id: %s\n User: %s\n Pass: <redacted>\n Host: %s\n Port: %s\n}\n"
	return fmt.Sprintf(format, c.confid, c.User, c.Host, c.Port)
}

// Global var to hold config
var conf *config

// Get-or-create a new configuration
func getConfig() (*config, error) {

	// Get a new config if one does not exist
	if conf == nil {

		log.Println("Creating new account config...")

		// Retrieve config from env
		newconf := &config{}
		err := envconfig.Process(envPrefix, newconf)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		newconf.confid = uuid.New()

		// Set this as the new config
		conf = newconf
	}

	return conf, nil
}

// Clear the config
func unsetConfig() {
	log.Println("Unsetting account config...")
	conf = nil
}

// Show usage / help
func configHelp() {
	log.Println(envconfig.Usage(envPrefix, &config{}))
}
