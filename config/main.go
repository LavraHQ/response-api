package config

import (
	"encoding/json"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// JwtConfig is the configuration stored in the JWT_SECRET env var.
type JwtConfig struct {
	Type   string
	Key    string
	Issuer string
}

// Decode the JwtConfig JSON object to the Struct
func (c *JwtConfig) Decode(value string) error {
	return json.Unmarshal([]byte(value), c)
}

// Specification is the specification of the env-based config.
type Specification struct {
	AuthJwtSecret JwtConfig `required:"true" split_words:"true"`
	DataURL       string    `required:"true" split_words:"true"`
	DataSecret    string    `required:"true" split_words:"true"`
}

// Config holds the config values.
var Config Specification

// ReadConfig reads the configuration from the environment
// variables.
func ReadConfig() error {
	err := envconfig.Process("response", &Config)
	if err != nil {
		fmt.Println("There was a problem loading the config.")
		fmt.Println(err.Error())

		return err
	}

	fmt.Println("Configuration has been read.")
	return nil
}
